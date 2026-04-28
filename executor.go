package polymarketcontracts

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethclient"
	"github.com/ivanzzeth/polymarket-go-contracts/sender"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

type contractCall struct {
	Target   common.Address
	Calldata []byte
	Value    *big.Int
}

type txExecutor struct {
	client      ethclient.EthClientInterface
	txSender    sender.TransactionSender
	getSafeAddr func(eoa common.Address) (common.Address, error)
	execSafeTx  func(safeSigner signer.SafeTradingSigner, chainID *big.Int, safeAddr, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int) (common.Hash, error)
}

func (e *txExecutor) executeEOA(call contractCall) (common.Hash, error) {
	txHash, err := e.txSender.SendEthereumTransaction(call.Target, call.Calldata, call.Value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send EOA transaction: %w", err)
	}
	return txHash, nil
}

func (e *txExecutor) executeSafe(safeSigner signer.SafeTradingSigner, chainID *big.Int, call contractCall) (common.Hash, error) {
	safeAddr, err := e.getSafeAddr(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}
	txHash, err := e.execSafeTx(safeSigner, chainID, safeAddr, call.Target, call.Value, call.Calldata, SafeOperationCall, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}
	return txHash, nil
}

func (e *txExecutor) executeBatchEOA(calls []contractCall) ([]common.Hash, error) {
	hashes := make([]common.Hash, 0, len(calls))
	for i, call := range calls {
		txHash, err := e.executeEOA(call)
		if err != nil {
			return hashes, fmt.Errorf("batch EOA tx %d failed: %w", i, err)
		}
		hashes = append(hashes, txHash)
	}
	return hashes, nil
}

func (e *txExecutor) executeBatchSafe(safeSigner signer.SafeTradingSigner, chainID *big.Int, calls []contractCall) ([]common.Hash, error) {
	hashes := make([]common.Hash, 0, len(calls))
	for i, call := range calls {
		txHash, err := e.executeSafe(safeSigner, chainID, call)
		if err != nil {
			return hashes, fmt.Errorf("batch Safe tx %d failed: %w", i, err)
		}
		hashes = append(hashes, txHash)

		if err := e.waitTxConfirmation(txHash, 3, 2*time.Minute); err != nil {
			return hashes, fmt.Errorf("batch Safe tx %d confirmation failed: %w", i, err)
		}
	}
	return hashes, nil
}

func (e *txExecutor) waitTxConfirmation(txHash common.Hash, confirmations uint64, timeout time.Duration) error {
	if e.client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("tx %s not confirmed after %v", txHash.Hex(), timeout)
		case <-ticker.C:
			receipt, err := e.client.TransactionReceipt(ctx, txHash)
			if err != nil || receipt == nil {
				continue
			}
			if confirmations == 0 {
				return nil
			}
			currentBlock, err := e.client.BlockNumber(ctx)
			if err != nil {
				continue
			}
			if currentBlock >= receipt.BlockNumber.Uint64()+confirmations {
				return nil
			}
		}
	}
}
