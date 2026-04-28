package polymarketcontracts

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethsig/eip712"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

// mockTransactionSender captures args for verification.
type mockTransactionSender struct {
	lastTo    common.Address
	lastData  []byte
	lastValue *big.Int
	calls     int
	retHash   common.Hash
	retErr    error
}

func (m *mockTransactionSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	m.lastTo = to
	m.lastData = data
	m.lastValue = value
	m.calls++
	return m.retHash, m.retErr
}

// mockSafeSigner implements signer.SafeTradingSigner for testing.
type mockSafeSigner struct {
	addr common.Address
}

func (m *mockSafeSigner) GetAddress() common.Address { return m.addr }
func (m *mockSafeSigner) SignTypedData(_ eip712.TypedData) ([]byte, error) {
	return []byte{0xAA}, nil
}
func (m *mockSafeSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	return common.Hash{}, nil
}

var _ signer.SafeTradingSigner = (*mockSafeSigner)(nil)

func TestExecuteEOA(t *testing.T) {
	expectedHash := common.HexToHash("0xabc123")
	mock := &mockTransactionSender{retHash: expectedHash}

	exec := &txExecutor{txSender: mock}
	call := contractCall{
		Target:   common.HexToAddress("0x1111"),
		Calldata: []byte{0x01, 0x02},
		Value:    big.NewInt(42),
	}

	hash, err := exec.executeEOA(call)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash != expectedHash {
		t.Errorf("expected hash %s, got %s", expectedHash.Hex(), hash.Hex())
	}
	if mock.lastTo != call.Target {
		t.Errorf("expected to %s, got %s", call.Target.Hex(), mock.lastTo.Hex())
	}
	if mock.lastValue.Cmp(call.Value) != 0 {
		t.Errorf("expected value %s, got %s", call.Value, mock.lastValue)
	}
	if string(mock.lastData) != string(call.Calldata) {
		t.Errorf("calldata mismatch")
	}
}

func TestExecuteEOA_Error(t *testing.T) {
	mock := &mockTransactionSender{retErr: errors.New("send failed")}
	exec := &txExecutor{txSender: mock}

	_, err := exec.executeEOA(contractCall{Value: big.NewInt(0)})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, mock.retErr) {
		t.Errorf("expected wrapped send error, got: %v", err)
	}
}

func TestExecuteSafe(t *testing.T) {
	expectedHash := common.HexToHash("0xdef456")
	safeAddr := common.HexToAddress("0xSafe")
	eoaAddr := common.HexToAddress("0xEOA")

	exec := &txExecutor{
		getSafeAddr: func(eoa common.Address) (common.Address, error) {
			if eoa != eoaAddr {
				t.Errorf("expected eoa %s, got %s", eoaAddr.Hex(), eoa.Hex())
			}
			return safeAddr, nil
		},
		execSafeTx: func(ss signer.SafeTradingSigner, chainID *big.Int, safe, to common.Address, value *big.Int, data []byte, op SafeOperation, gas *big.Int) (common.Hash, error) {
			if safe != safeAddr {
				t.Errorf("expected safe %s, got %s", safeAddr.Hex(), safe.Hex())
			}
			return expectedHash, nil
		},
	}

	ms := &mockSafeSigner{addr: eoaAddr}
	hash, err := exec.executeSafe(ms, big.NewInt(137), contractCall{
		Target:   common.HexToAddress("0xTarget"),
		Calldata: []byte{0xAB},
		Value:    big.NewInt(0),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash != expectedHash {
		t.Errorf("expected hash %s, got %s", expectedHash.Hex(), hash.Hex())
	}
}

func TestExecuteSafe_GetSafeAddrError(t *testing.T) {
	exec := &txExecutor{
		getSafeAddr: func(eoa common.Address) (common.Address, error) {
			return common.Address{}, errors.New("no safe")
		},
	}

	ms := &mockSafeSigner{addr: common.HexToAddress("0xEOA")}
	_, err := exec.executeSafe(ms, big.NewInt(137), contractCall{Value: big.NewInt(0)})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestExecuteBatchEOA(t *testing.T) {
	hash1 := common.HexToHash("0x0001")
	hash2 := common.HexToHash("0x0002")
	callIdx := 0

	exec := &txExecutor{
		txSender: &sequentialMockSender{hashes: []common.Hash{hash1, hash2}, idx: &callIdx},
	}

	calls := []contractCall{
		{Target: common.HexToAddress("0xA"), Calldata: []byte{1}, Value: big.NewInt(0)},
		{Target: common.HexToAddress("0xB"), Calldata: []byte{2}, Value: big.NewInt(0)},
	}

	result, err := exec.executeBatchEOA(calls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 hashes, got %d", len(result))
	}
	if result[0] != hash1 || result[1] != hash2 {
		t.Errorf("hash mismatch: got %v", result)
	}
}

func TestExecuteBatchEOA_PartialFailure(t *testing.T) {
	callIdx := 0
	exec := &txExecutor{
		txSender: &failOnSecondSender{idx: &callIdx},
	}

	calls := []contractCall{
		{Target: common.HexToAddress("0xA"), Calldata: []byte{1}, Value: big.NewInt(0)},
		{Target: common.HexToAddress("0xB"), Calldata: []byte{2}, Value: big.NewInt(0)},
	}

	result, err := exec.executeBatchEOA(calls)
	if err == nil {
		t.Fatal("expected error on second call")
	}
	if len(result) != 1 {
		t.Errorf("expected 1 successful hash before failure, got %d", len(result))
	}
}

// sequentialMockSender returns pre-defined hashes in order.
type sequentialMockSender struct {
	hashes []common.Hash
	idx    *int
}

func (s *sequentialMockSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	i := *s.idx
	*s.idx++
	if i < len(s.hashes) {
		return s.hashes[i], nil
	}
	return common.Hash{}, errors.New("no more hashes")
}

// failOnSecondSender succeeds on first call, fails on second.
type failOnSecondSender struct {
	idx *int
}

func (s *failOnSecondSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	i := *s.idx
	*s.idx++
	if i == 0 {
		return common.HexToHash("0x0001"), nil
	}
	return common.Hash{}, errors.New("boom")
}
