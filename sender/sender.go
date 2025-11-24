package sender

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TransactionSender defines the interface for sending Ethereum transactions
type TransactionSender interface {
	SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error)
}
