package polymarketcontracts

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBuildSafeTransactionTypedData(t *testing.T) {
	chainID := big.NewInt(137)
	safeAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
	to := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	value := big.NewInt(0)
	data := []byte{0x01, 0x02, 0x03}
	operation := SafeOperationCall
	safeTxGas := big.NewInt(100000)
	baseGas := big.NewInt(50000)
	gasPrice := big.NewInt(1000000000)
	gasToken := common.HexToAddress("0x0000000000000000000000000000000000000000")
	refundReceiver := common.HexToAddress("0x0000000000000000000000000000000000000000")
	nonce := big.NewInt(0)

	typedData := BuildSafeTransactionTypedData(
		chainID, safeAddr, to, value, data, operation,
		safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce,
	)

	// Verify basic structure
	if typedData.PrimaryType != "SafeTx" {
		t.Errorf("Expected PrimaryType 'SafeTx', got '%s'", typedData.PrimaryType)
	}

	// Verify domain
	if typedData.Domain.ChainId != "137" {
		t.Errorf("Expected ChainId '137', got '%s'", typedData.Domain.ChainId)
	}
	if typedData.Domain.VerifyingContract != safeAddr.Hex() {
		t.Errorf("Expected VerifyingContract '%s', got '%s'", safeAddr.Hex(), typedData.Domain.VerifyingContract)
	}

	// Verify types
	if len(typedData.Types["EIP712Domain"]) != 2 {
		t.Errorf("Expected 2 EIP712Domain fields, got %d", len(typedData.Types["EIP712Domain"]))
	}
	if len(typedData.Types["SafeTx"]) != 10 {
		t.Errorf("Expected 10 SafeTx fields, got %d", len(typedData.Types["SafeTx"]))
	}

	// Verify message fields
	if typedData.Message["to"] != to.Hex() {
		t.Errorf("Expected to '%s', got '%s'", to.Hex(), typedData.Message["to"])
	}
	if typedData.Message["value"] != "0" {
		t.Errorf("Expected value '0', got '%s'", typedData.Message["value"])
	}
	if typedData.Message["data"] != "0x010203" {
		t.Errorf("Expected data '0x010203', got '%s'", typedData.Message["data"])
	}
	if typedData.Message["operation"] != "0" {
		t.Errorf("Expected operation '0', got '%s'", typedData.Message["operation"])
	}
	if typedData.Message["safeTxGas"] != "100000" {
		t.Errorf("Expected safeTxGas '100000', got '%s'", typedData.Message["safeTxGas"])
	}
	if typedData.Message["nonce"] != "0" {
		t.Errorf("Expected nonce '0', got '%s'", typedData.Message["nonce"])
	}
}

func TestBuildSafeTransactionTypedData_DelegateCall(t *testing.T) {
	chainID := big.NewInt(137)
	safeAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
	to := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	value := big.NewInt(1000000)
	data := []byte{}
	operation := SafeOperationDelegateCall
	safeTxGas := big.NewInt(200000)
	baseGas := big.NewInt(100000)
	gasPrice := big.NewInt(2000000000)
	gasToken := common.HexToAddress("0x0000000000000000000000000000000000000000")
	refundReceiver := common.HexToAddress("0x9876543210987654321098765432109876543210")
	nonce := big.NewInt(5)

	typedData := BuildSafeTransactionTypedData(
		chainID, safeAddr, to, value, data, operation,
		safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce,
	)

	// Verify operation is DelegateCall (1)
	if typedData.Message["operation"] != "1" {
		t.Errorf("Expected operation '1', got '%s'", typedData.Message["operation"])
	}

	// Verify value
	if typedData.Message["value"] != "1000000" {
		t.Errorf("Expected value '1000000', got '%s'", typedData.Message["value"])
	}

	// Verify empty data
	if typedData.Message["data"] != "0x" {
		t.Errorf("Expected data '0x', got '%s'", typedData.Message["data"])
	}

	// Verify nonce
	if typedData.Message["nonce"] != "5" {
		t.Errorf("Expected nonce '5', got '%s'", typedData.Message["nonce"])
	}

	// Verify refundReceiver
	if typedData.Message["refundReceiver"] != refundReceiver.Hex() {
		t.Errorf("Expected refundReceiver '%s', got '%s'", refundReceiver.Hex(), typedData.Message["refundReceiver"])
	}
}
