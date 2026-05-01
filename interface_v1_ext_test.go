package polymarketcontracts

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	conditional_tokens "github.com/ivanzzeth/polymarket-go-contracts/v2/contracts/conditional-tokens"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/v2/contracts/ctf-collateral-adapter"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/v2/contracts/neg-risk-ctf-collateral-adapter"
	"github.com/ivanzzeth/polymarket-go-contracts/v2/signer"

	"github.com/ivanzzeth/ethsig/eip712"
)

// mockEOASigner satisfies signer.EOATradingSigner for testing.
type mockEOASigner struct {
	addr common.Address
}

func (m *mockEOASigner) GetAddress() common.Address { return m.addr }
func (m *mockEOASigner) SignTypedData(_ eip712.TypedData) ([]byte, error) {
	return []byte{0xBB}, nil
}
func (m *mockEOASigner) SignTransactionWithChainID(tx *types.Transaction, _ *big.Int) (*types.Transaction, error) {
	return tx, nil
}

var _ signer.EOATradingSigner = (*mockEOASigner)(nil)

// capturingMockSender records all target addresses across calls.
type capturingMockSender struct {
	targets *[]common.Address
}

func (c *capturingMockSender) SendEthereumTransaction(to common.Address, _ []byte, _ *big.Int) (common.Hash, error) {
	*c.targets = append(*c.targets, to)
	return common.BigToHash(big.NewInt(int64(len(*c.targets)))), nil
}

func newV1ExtTestCI() (*ContractInterface, *mockTransactionSender) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV1EXT")}
	ci := &ContractInterface{
		contractConfig: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}
	return ci, mock
}

func TestSplitPositionWithCollateral_USDCE_MatchesV1(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.SplitPositionForEOA(ctx, nil, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("V1 SplitPositionForEOA: %v", err)
	}
	v1Data := make([]byte, len(mock.lastData))
	copy(v1Data, mock.lastData)
	v1Target := mock.lastTo

	_, err = ci.SplitPositionWithCollateralForEOA(ctx, nil, CollateralUSDCE, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("WithCollateral(USDCE): %v", err)
	}

	if !bytes.Equal(v1Data, mock.lastData) {
		t.Error("WithCollateral(USDCE) calldata differs from V1")
	}
	if v1Target != mock.lastTo {
		t.Error("WithCollateral(USDCE) target differs from V1")
	}
}

func TestSplitPositionWithCollateral_PUSD_UsesAdapter(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.SplitPositionWithCollateralForEOA(ctx, nil, CollateralPUSD, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("WithCollateral(PUSD): %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter %s, got %s",
			MATIC_CONTRACTS.CtfCollateralAdapter.Hex(), mock.lastTo.Hex())
	}

	parsedABI, _ := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(mock.lastData, expected) {
		t.Error("PUSD split calldata mismatch")
	}
}

func TestMergePositionsWithCollateral_PUSD_UsesAdapter(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.MergePositionsWithCollateralForEOA(ctx, nil, CollateralPUSD, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("MergeWithCollateral(PUSD): %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestRedeemPositionsWithCollateral_PUSD_UsesAdapter(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}

	_, err := ci.RedeemPositionsWithCollateralForEOA(ctx, nil, CollateralPUSD, testCondID, indexSets)
	if err != nil {
		t.Fatalf("RedeemWithCollateral(PUSD): %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestSplitNegRiskWithCollateral_PUSD_UsesNRAdapter(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.SplitPositionNegRiskWithCollateralForEOA(ctx, nil, CollateralPUSD, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("SplitNegRisk(PUSD): %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.NegRiskCtfCollateralAdapter {
		t.Errorf("expected target NegRiskCtfCollateralAdapter %s, got %s",
			MATIC_CONTRACTS.NegRiskCtfCollateralAdapter.Hex(), mock.lastTo.Hex())
	}

	parsedABI, _ := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(mock.lastData, expected) {
		t.Error("NegRisk PUSD split calldata mismatch")
	}
}

func TestPUSDConvenienceMethod_DelegatesToWithCollateral(t *testing.T) {
	ci, mock := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.SplitPositionPUSDForEOA(ctx, nil, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("SplitPositionPUSDForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("PUSD convenience should target CtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestWrapCollateralForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xWRAP")}
	eoaAddr := common.HexToAddress("0xEOA1")
	eoaSigner := &mockEOASigner{addr: eoaAddr}
	ci := &ContractInterface{
		contractConfig: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}

	asset := MATIC_CONTRACTS.Collateral
	amount := big.NewInt(5000000)

	_, err := ci.WrapCollateralForEOA(context.Background(), eoaSigner, asset, amount)
	if err != nil {
		t.Fatalf("WrapCollateralForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CollateralOnramp {
		t.Errorf("expected target CollateralOnramp %s, got %s",
			MATIC_CONTRACTS.CollateralOnramp.Hex(), mock.lastTo.Hex())
	}
}

func TestUnwrapCollateralForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xUNWRAP")}
	eoaAddr := common.HexToAddress("0xEOA2")
	eoaSigner := &mockEOASigner{addr: eoaAddr}
	ci := &ContractInterface{
		contractConfig: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}

	asset := MATIC_CONTRACTS.Collateral
	amount := big.NewInt(3000000)

	_, err := ci.UnwrapCollateralForEOA(context.Background(), eoaSigner, asset, amount)
	if err != nil {
		t.Fatalf("UnwrapCollateralForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CollateralOfframp {
		t.Errorf("expected target CollateralOfframp %s, got %s",
			MATIC_CONTRACTS.CollateralOfframp.Hex(), mock.lastTo.Hex())
	}
}

func TestEnableTradingV2ForEOA(t *testing.T) {
	callIdx := 0
	hashes := make([]common.Hash, 7)
	for i := range hashes {
		hashes[i] = common.BigToHash(big.NewInt(int64(i + 1)))
	}
	mock := &sequentialMockSender{hashes: hashes, idx: &callIdx}
	ci := &ContractInterface{
		contractConfig: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}

	result, err := ci.EnableTradingV2ForEOA(context.Background(), nil)
	if err != nil {
		t.Fatalf("EnableTradingV2ForEOA: %v", err)
	}

	if len(result) != 7 {
		t.Fatalf("expected 7 tx hashes (3 pUSD approvals + 4 CTF approvals), got %d", len(result))
	}
}

func TestEnableTradingV2ForEOA_CTFApprovalTargets(t *testing.T) {
	var targets []common.Address
	mock := &capturingMockSender{targets: &targets}
	ci := &ContractInterface{
		contractConfig: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}

	_, err := ci.EnableTradingV2ForEOA(context.Background(), nil)
	if err != nil {
		t.Fatalf("EnableTradingV2ForEOA: %v", err)
	}

	if len(targets) != 7 {
		t.Fatalf("expected 7 calls, got %d", len(targets))
	}

	for i := 0; i < 3; i++ {
		if targets[i] != MATIC_CONTRACTS.CollateralToken {
			t.Errorf("call %d: expected target CollateralToken %s, got %s",
				i, MATIC_CONTRACTS.CollateralToken.Hex(), targets[i].Hex())
		}
	}

	for i := 3; i < 7; i++ {
		if targets[i] != MATIC_CONTRACTS.ConditionalTokens {
			t.Errorf("call %d: expected target ConditionalTokens %s, got %s",
				i, MATIC_CONTRACTS.ConditionalTokens.Hex(), targets[i].Hex())
		}
	}
}

func TestWithCollateral_InvalidType(t *testing.T) {
	ci, _ := newV1ExtTestCI()
	ctx := context.Background()

	_, err := ci.SplitPositionWithCollateralForEOA(ctx, nil, CollateralType(99), testCondID, testPartition, testAmount)
	if err == nil {
		t.Error("expected error for invalid collateral type")
	}
}

func TestV1SplitPositionForEOA_UsesCTFDirectly(t *testing.T) {
	ci, mock := newV1ExtTestCI()

	_, err := ci.SplitPositionForEOA(context.Background(), nil, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("SplitPositionForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.ConditionalTokens {
		t.Errorf("V1 split should target ConditionalTokens %s, got %s",
			MATIC_CONTRACTS.ConditionalTokens.Hex(), mock.lastTo.Hex())
	}

	parsedABI, _ := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", MATIC_CONTRACTS.Collateral, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(mock.lastData, expected) {
		t.Error("V1 split calldata regression")
	}
}
