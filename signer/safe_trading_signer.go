package signer

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethsig"
	"github.com/ivanzzeth/ethsig/eip712"
	"github.com/ivanzzeth/polymarket-go-contracts/sender"
)

// SafeTradingSigner defines the interface requirements for Safe trading operations
type SafeTradingSigner interface {
	ethsig.AddressGetter
	ethsig.TypedDataSigner
	sender.TransactionSender
}

// SimpleSafeTradingSigner is a generic implementation of SafeTradingSigner
type SimpleSafeTradingSigner struct {
	addr            common.Address
	typedDataSigner ethsig.TypedDataSigner
	txSender        sender.TransactionSender
}

// NewSimpleSafeTradingSigner creates a new SimpleSafeTradingSigner
func NewSimpleSafeTradingSigner(addr common.Address, typedDataSigner ethsig.TypedDataSigner, txSender sender.TransactionSender) *SimpleSafeTradingSigner {
	return &SimpleSafeTradingSigner{
		addr:            addr,
		typedDataSigner: typedDataSigner,
		txSender:        txSender,
	}
}

// GetAddress returns the signer's address
func (s *SimpleSafeTradingSigner) GetAddress() common.Address {
	return s.addr
}

// SignTypedData signs EIP-712 typed data
func (s *SimpleSafeTradingSigner) SignTypedData(typedData eip712.TypedData) ([]byte, error) {
	return s.typedDataSigner.SignTypedData(typedData)
}

// SendEthereumTransaction sends an Ethereum transaction
func (s *SimpleSafeTradingSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	return s.txSender.SendEthereumTransaction(to, data, value)
}

// SafeTradingSingleMpcSigner is a SafeTradingSigner implementation using Cobo MPC
type SafeTradingSingleMpcSigner struct {
	*SimpleSafeTradingSigner
	chainId *big.Int
	client  bind.ContractBackend
}

func NewSafeTradingSingleMpcSigner(chainId *big.Int, client bind.ContractBackend, signer *CoboMpcSigner) (*SafeTradingSingleMpcSigner, error) {
	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return &SafeTradingSingleMpcSigner{
		SimpleSafeTradingSigner: NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender),
		chainId:                 chainId,
		client:                  client,
	}, nil
}

// NewSafeTradingPrivateKeySigner creates a SafeTradingSigner from a private key
func NewSafeTradingPrivateKeySigner(chainId *big.Int, client bind.ContractBackend, privateKey *ecdsa.PrivateKey) (*SimpleSafeTradingSigner, error) {
	signer := ethsig.NewEthPrivateKeySigner(privateKey)
	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender), nil
}

// NewSafeTradingPrivateKeySignerFromHex creates a SafeTradingSigner from a private key hex string
func NewSafeTradingPrivateKeySignerFromHex(chainId *big.Int, client bind.ContractBackend, privateKeyHex string) (*SimpleSafeTradingSigner, error) {
	signer, err := ethsig.NewEthPrivateKeySignerFromPrivateKeyHex(privateKeyHex)
	if err != nil {
		return nil, err
	}

	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender), nil
}

// NewSafeTradingKeystoreSignerFromPath creates a SafeTradingSigner from a keystore file or directory
func NewSafeTradingKeystoreSignerFromPath(chainId *big.Int, client bind.ContractBackend, keystorePath string, address common.Address, password string, scryptConfig *ethsig.KeystoreScryptConfig) (*SimpleSafeTradingSigner, error) {
	signer, err := ethsig.NewKeystoreSignerFromPath(keystorePath, address, password, scryptConfig)
	if err != nil {
		return nil, err
	}

	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender), nil
}

// NewSafeTradingKeystoreSignerFromFile creates a SafeTradingSigner from a keystore file
func NewSafeTradingKeystoreSignerFromFile(chainId *big.Int, client bind.ContractBackend, keystoreFile string, address common.Address, password string, scryptConfig *ethsig.KeystoreScryptConfig) (*SimpleSafeTradingSigner, error) {
	signer, err := ethsig.NewKeystoreSignerFromFile(keystoreFile, address, password, scryptConfig)
	if err != nil {
		return nil, err
	}

	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender), nil
}

// NewSafeTradingKeystoreSignerFromDirectory creates a SafeTradingSigner from a keystore directory
func NewSafeTradingKeystoreSignerFromDirectory(chainId *big.Int, client bind.ContractBackend, keystoreDir string, address common.Address, password string, scryptConfig *ethsig.KeystoreScryptConfig) (*SimpleSafeTradingSigner, error) {
	signer, err := ethsig.NewKeystoreSignerFromDirectory(keystoreDir, address, password, scryptConfig)
	if err != nil {
		return nil, err
	}

	txSender, err := GetTransactionSenderBySigner(chainId, client, signer)
	if err != nil {
		return nil, err
	}

	return NewSimpleSafeTradingSigner(signer.GetAddress(), signer, txSender), nil
}
