package signer

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
	"github.com/ivanzzeth/ethsig/eip712"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/sirupsen/logrus"
)

type CoboMpcSigner struct {
	address common.Address

	coboPrivateKey string
	coboClient     *coboWaas2.APIClient
	coboChainId    string
	ethChainId     *big.Int
	coboWalletId   string
	coboEnv        int

	maxRetryCount int

	// logger *logrus.Logger
}

func NewCoboMpcSigner(env int, privateKey, coboChainId string, ethChainId *big.Int, walletId string, address common.Address) (*CoboMpcSigner, error) {
	if env != coboWaas2.ProdEnv && env != coboWaas2.DevEnv {
		return nil, fmt.Errorf("env should be coboWaas2.ProdEnv or coboWaas2.DevEnv, got %d", env)
	}

	signer := CoboMpcSigner{coboEnv: env, coboPrivateKey: privateKey, coboChainId: coboChainId, ethChainId: ethChainId, coboWalletId: walletId, maxRetryCount: 100, address: address}

	configuration := coboWaas2.NewConfiguration()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	configuration.HTTPClient = client
	apiClient := coboWaas2.NewAPIClient(configuration)

	signer.coboClient = apiClient

	// logger := logrus.New()
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp:   true,
	// 	TimestampFormat: time.StampMilli,
	// })
	// signer.logger = logger

	return &signer, nil
}

func (s CoboMpcSigner) GetAddress() common.Address {
	return s.address
}

func (s CoboMpcSigner) CoboChainId() string {
	return s.coboChainId
}

// SignTypedData signs EIP-712 typed data
func (s *CoboMpcSigner) SignTypedData(typedData eip712.TypedData) ([]byte, error) {
	mpcSource := coboWaas2.NewMpcMessageSignSource(coboWaas2.MESSAGESIGNSOURCETYPE_ORG_CONTROLLED, s.coboWalletId, s.address.Hex())
	signSource := coboWaas2.MpcMessageSignSourceAsMessageSignSource(mpcSource)

	// Convert TypedData to map[string]interface{} for Cobo API
	structuredData := map[string]interface{}{
		"types":       typedData.Types,
		"primaryType": typedData.PrimaryType,
		"domain":      typedData.Domain.Map(),
		"message":     typedData.Message,
	}

	signDestination := coboWaas2.NewEvmEIP712MessageSignDestination(coboWaas2.MESSAGESIGNDESTINATIONTYPE_EVM_EIP_712_SIGNATURE, structuredData)
	destination := coboWaas2.EvmEIP712MessageSignDestinationAsMessageSignDestination(signDestination)

	messageSignParams := *coboWaas2.NewMessageSignParams(
		s.createRequestId(),
		s.coboChainId,
		signSource,
		destination,
	)

	resp, r, err := s.coboClient.TransactionsAPI.CreateMessageSignTransaction(s.getCtx()).MessageSignParams(messageSignParams).Execute()
	formattedResp, err := s.formatResponse(resp, r, err)
	if err != nil {
		return nil, err
	}

	txDetail, err := s.WaitTransactionStatus(formattedResp.TransactionId, coboWaas2.TRANSACTIONSTATUS_COMPLETED, s.maxRetryCount)
	if err != nil {
		return nil, err
	}

	// Extract signature from transaction result
	if txDetail.Result == nil {
		return nil, fmt.Errorf("transaction has no result")
	}

	if txDetail.Result.TransactionSignatureResult == nil {
		return nil, fmt.Errorf("transaction result has no signature result")
	}

	sigStr := txDetail.Result.TransactionSignatureResult.Signature
	// Remove 0x prefix if present for processing
	if len(sigStr) > 2 && sigStr[:2] == "0x" {
		sigStr = sigStr[2:]
	}

	// Decode and validate signature
	signature := common.FromHex(sigStr)
	if len(signature) != 65 {
		return nil, fmt.Errorf("invalid signature length: expected 65, got %d", len(signature))
	}

	// Ensure v value is in the correct format (27/28)
	if signature[64] < 27 {
		signature[64] += 27
	}

	// Return signature as bytes
	return signature, nil
}

func (m *CoboMpcSigner) CallContract(to, callData, value string, fee *coboWaas2.TransactionRequestFee) (*coboWaas2.CreateTransferTransaction201Response, error) {
	//source
	mpcSource := coboWaas2.NewMpcContractCallSource(
		coboWaas2.CONTRACTCALLSOURCETYPE_ORG_CONTROLLED,
		m.coboWalletId,
		m.address.Hex(),
	)
	source := coboWaas2.MpcContractCallSourceAsContractCallSource(mpcSource)
	//des
	evmContractCallDes := coboWaas2.NewEvmContractCallDestination(coboWaas2.CONTRACTCALLDESTINATIONTYPE_EVM_CONTRACT, to, callData)
	if value != "" && value != "0" {
		evmContractCallDes.SetValue(value)
	}
	des := coboWaas2.EvmContractCallDestinationAsContractCallDestination(evmContractCallDes)
	//param
	contractCallParams := *coboWaas2.NewContractCallParams(
		m.createRequestId(),
		m.coboChainId,
		source,
		des,
	)
	if fee != nil {
		contractCallParams.Fee = fee
	}

	resp, r, err := m.coboClient.TransactionsAPI.CreateContractCallTransaction(m.getCtx()).ContractCallParams(contractCallParams).Execute()
	return m.formatResponse(resp, r, err)
}

func (m *CoboMpcSigner) WaitTransactionStatus(transactionId string, status coboWaas2.TransactionStatus, maxTryTime int) (*coboWaas2.TransactionDetail, error) {
	if status == coboWaas2.TRANSACTIONSTATUS_FAILED || status == coboWaas2.TRANSACTIONSTATUS_REJECTED {
		return nil, fmt.Errorf("invalid status")
	}

	tryTime := 1
	for {
		if tryTime > maxTryTime {
			return nil, fmt.Errorf("wait transaction get max try times")
		}
		resp, err := m.GetTransactionByTransactionId(transactionId)
		if err != nil {
			return nil, err
		}

		switch resp.Status {
		case status:
			//done
			return resp, nil
		case coboWaas2.TRANSACTIONSTATUS_REJECTED:
			return nil, fmt.Errorf("transaction id: %s, request id: %s, cobo id: %s, REJECTED", transactionId, *resp.RequestId, *resp.CoboId)
		case coboWaas2.TRANSACTIONSTATUS_FAILED:
			return nil, fmt.Errorf("transaction id: %s, request id: %s, cobo id: %s, FAILED", transactionId, *resp.RequestId, *resp.CoboId)
		default:
			//continue when got other status
			log.Printf("transaction id: %s, request id: %s, cobo id: %s, status: %s", transactionId, *resp.RequestId, *resp.CoboId, resp.Status)
		}

		time.Sleep(3 * time.Second)
		tryTime++
	}
}

func (m *CoboMpcSigner) GetTransactionByTransactionId(transactionId string) (*coboWaas2.TransactionDetail, error) {
	resp, _, err := m.coboClient.TransactionsAPI.GetTransactionById(m.getCtx(), transactionId).Execute()
	return resp, err
}

func (s *CoboMpcSigner) createRequestId() string {
	return fmt.Sprintf("cobo-mpc-go-v2-%d", time.Now().UnixMilli())
}

func (m *CoboMpcSigner) getCtx() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, coboWaas2.ContextEnv, m.coboEnv)
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: m.coboPrivateKey,
	})
	return ctx
}

func (m *CoboMpcSigner) formatResponse(resp *coboWaas2.CreateTransferTransaction201Response, r *http.Response, err error) (*coboWaas2.CreateTransferTransaction201Response, error) {
	if err == nil {
		return resp, nil
	}

	if r != nil && r.Body != nil {
		defer r.Body.Close()
		body, readErr := io.ReadAll(r.Body)
		if readErr != nil {
			return nil, fmt.Errorf("err: %s, unable to read response body: %s", err.Error(), readErr)
		}

		return nil, fmt.Errorf("err: %s, body: %s", err.Error(), string(body))
	}

	return nil, fmt.Errorf("err: %s, nil body", err.Error())
}
