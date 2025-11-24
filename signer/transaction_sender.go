package signer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ivanzzeth/ethsig"
	"github.com/ivanzzeth/polymarket-go-contracts/sender"
)

// GetTransactionSenderBySigner creates a TransactionSender based on the signer type
func GetTransactionSenderBySigner(chainId *big.Int, client bind.ContractBackend, signerInstance any) (sender.TransactionSender, error) {
	switch s := signerInstance.(type) {
	case *CoboMpcSigner:
		return GetTransactionSenderByCoboMpcTransactionSender(client, s)
	case TransactionSignerAndAddrGetter:
		return GetTransactionSenderByTransactionSignerAndAddrGetter(chainId, client, s)
	default:
		return nil, fmt.Errorf("unsupported signer type %T", signerInstance)
	}
}

// TransactionSignerAndAddrGetter combines transaction signing and address retrieval
type TransactionSignerAndAddrGetter interface {
	ethsig.AddressGetter
	ethsig.TransactionSigner
}

// TransactionSenderByTransactionSigner implements TransactionSender using a transaction signer
type TransactionSenderByTransactionSigner struct {
	chainId  *big.Int
	client   bind.ContractBackend
	txSigner TransactionSignerAndAddrGetter
}

// GetTransactionSenderByTransactionSignerAndAddrGetter creates a TransactionSender from a transaction signer
func GetTransactionSenderByTransactionSignerAndAddrGetter(chainId *big.Int, client bind.ContractBackend, txSigner TransactionSignerAndAddrGetter) (sender.TransactionSender, error) {
	return &TransactionSenderByTransactionSigner{chainId: chainId, client: client, txSigner: txSigner}, nil
}

// SendEthereumTransaction sends an Ethereum transaction using the transaction signer
func (s *TransactionSenderByTransactionSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	ctx := context.Background()

	// Get nonce
	nonce, err := s.client.PendingNonceAt(ctx, s.txSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     s.txSigner.GetAddress(),
		To:       &to,
		Value:    value,
		Data:     data,
		GasPrice: gasPrice,
	}
	gasLimit, err := s.client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Create and sign the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := s.txSigner.SignTransactionWithChainID(tx, s.chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the signed transaction
	err = s.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash(), nil
}

// CoboMpcTransactionSender implements TransactionSender using Cobo MPC wallet
type CoboMpcTransactionSender struct {
	client bind.ContractBackend
	signer *CoboMpcSigner
}

// GetTransactionSenderByCoboMpcTransactionSender creates a TransactionSender for Cobo MPC
func GetTransactionSenderByCoboMpcTransactionSender(client bind.ContractBackend, mpcSigner *CoboMpcSigner) (sender.TransactionSender, error) {
	return &CoboMpcTransactionSender{client: client, signer: mpcSigner}, nil
}

// SendEthereumTransaction sends an Ethereum transaction using Cobo MPC wallet
func (s *CoboMpcTransactionSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Hash{}, err
	}

	// Increase gas price by 50%
	gasPrice.Mul(gasPrice, big.NewInt(15))
	gasPrice.Div(gasPrice, big.NewInt(10))

	gasLimit, err := s.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  s.signer.GetAddress(),
		To:    &to,
		Data:  data,
		Value: value,
	})
	if err != nil {
		return common.Hash{}, err
	}

	fee := cobo_waas2.NewTransactionRequestEvmLegacyFee(gasPrice.String(), cobo_waas2.FEETYPE_EVM_LEGACY, s.signer.CoboChainId())
	fee.SetGasLimit(big.NewInt(int64(gasLimit)).String())

	paramFee := cobo_waas2.TransactionRequestEvmLegacyFeeAsTransactionRequestFee(fee)

	txResp, err := s.signer.CallContract(to.Hex(), fmt.Sprintf("0x%x", data), value.String(), &paramFee)
	if err != nil {
		return common.Hash{}, err
	}

	txDetail, err := s.signer.WaitTransactionStatus(txResp.TransactionId, cobo_waas2.TRANSACTIONSTATUS_CONFIRMING, 100)
	if err != nil {
		return common.Hash{}, err
	}

	txHash := common.HexToHash(*txDetail.TransactionHash)
	return txHash, nil
}
