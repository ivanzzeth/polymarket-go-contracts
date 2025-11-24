package signer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethsig"
)

// SafeTradingSigner defines the interface requirements for Safe trading operations
type SafeTradingSigner interface {
	ethsig.AddressGetter
	ethsig.TypedDataSigner
	TransactionSender
}

type SafeTradingSingleMpcSigner struct {
	chainId *big.Int
	client  bind.ContractBackend
	*CoboMpcSigner
}

func NewSafeTradingSingleMpcSigner(chainId *big.Int, client bind.ContractBackend, signer *CoboMpcSigner) (*SafeTradingSingleMpcSigner, error) {
	return &SafeTradingSingleMpcSigner{
		chainId:       chainId,
		client:        client,
		CoboMpcSigner: signer,
	}, nil
}

func (s *SafeTradingSingleMpcSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	txSender, err := GetTransactionSenderBySigner(s.chainId, s.client, s.CoboMpcSigner)
	if err != nil {
		return common.Hash{}, err
	}

	return txSender.SendEthereumTransaction(to, data, value)
}
