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
)

type TransactionSender interface {
	SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error)
}

func GetTransactionSenderBySigner(chainId *big.Int, client bind.ContractBackend, signer any) (TransactionSender, error) {
	switch s := signer.(type) {
	case *CoboMpcSigner:
		return GetTransactionSenderByCoboMpcTransactionSender(client, s)
	case TransactionSignerAndAddrGetter:
		return GetTransactionSenderByCoboMpcTransactionSignerAndAddrGetter(chainId, client, s)
	default:
		return nil, fmt.Errorf("unsupported signer type %T", signer)
	}
}

type TransactionSenderrByTransactionSigner struct {
	chainId  *big.Int
	client   bind.ContractBackend
	txSigner TransactionSignerAndAddrGetter
}

func GetTransactionSenderByCoboMpcTransactionSignerAndAddrGetter(chainId *big.Int, client bind.ContractBackend, signer TransactionSignerAndAddrGetter) (TransactionSender, error) {
	return &TransactionSenderrByTransactionSigner{chainId: chainId, client: client, txSigner: signer}, nil
}

type TransactionSignerAndAddrGetter interface {
	ethsig.AddressGetter
	ethsig.TransactionSigner
}

func (b *TransactionSenderrByTransactionSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	ctx := context.Background()

	// Get nonce
	nonce, err := b.client.PendingNonceAt(ctx, b.txSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     b.txSigner.GetAddress(),
		To:       &to,
		Value:    value,
		Data:     data,
		GasPrice: gasPrice,
	}
	gasLimit, err := b.client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Create and sign the transaction using BoundContract
	// This is a workaround since we need to create a raw transaction
	// We'll use the ethereum types directly
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := b.txSigner.SignTransactionWithChainID(tx, b.chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the signed transaction
	err = b.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash(), nil
}

func GetTransactionSenderByCoboMpcTransactionSender(client bind.ContractBackend, mpcSigner *CoboMpcSigner) (TransactionSender, error) {
	return &CoboMpcTransactionSender{client: client, signer: mpcSigner}, nil
}

type CoboMpcTransactionSender struct {
	client bind.ContractBackend
	signer *CoboMpcSigner
}

func (s *CoboMpcTransactionSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.Hash{}, err
	}

	gasPrice.Mul(gasPrice, big.NewInt(15))
	gasPrice.Div(gasPrice, big.NewInt(10))
	// fee := &cobo_waas2.TransactionRequestFee{
	// 	TransactionRequestEvmLegacyFee: cobo_waas2.NewTransactionRequestEvmLegacyFee(gasPrice.String(), cobo_waas2.FEETYPE_EVM_LEGACY, s.signer.coboChainId),
	// }
	gasLimit, err := s.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  s.signer.GetAddress(),
		To:    &to,
		Data:  data,
		Value: value,
	})
	if err != nil {
		return common.Hash{}, err
	}

	fee := cobo_waas2.NewTransactionRequestEvmLegacyFee(gasPrice.String(), cobo_waas2.FEETYPE_EVM_LEGACY, s.signer.coboChainId)
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
