package polymarketcontracts

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/ctf-collateral-adapter"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-ctf-collateral-adapter"
)

func newV2TestInstance(mock *mockTransactionSender) *ContractInterfaceV2 {
	return &ContractInterfaceV2{
		config: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
	}
}

func TestV2EnableTradingForEOA_CallCount(t *testing.T) {
	callIdx := 0
	hashes := make([]common.Hash, 7)
	for i := range hashes {
		hashes[i] = common.BigToHash(big.NewInt(int64(i + 1)))
	}
	v2 := &ContractInterfaceV2{
		config: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: &sequentialMockSender{hashes: hashes, idx: &callIdx},
		},
	}

	result, err := v2.EnableTradingForEOA(context.Background())
	if err != nil {
		t.Fatalf("EnableTradingForEOA: %v", err)
	}

	if len(result) != 7 {
		t.Fatalf("expected 7 tx hashes, got %d", len(result))
	}
}

func TestV2EnableTradingForEOA_Targets(t *testing.T) {
	var targets []common.Address
	v2 := &ContractInterfaceV2{
		config: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: &capturingMockSender{targets: &targets},
		},
	}

	_, err := v2.EnableTradingForEOA(context.Background())
	if err != nil {
		t.Fatalf("EnableTradingForEOA: %v", err)
	}

	if len(targets) != 7 {
		t.Fatalf("expected 7 calls, got %d", len(targets))
	}

	for i := 0; i < 3; i++ {
		if targets[i] != MATIC_CONTRACTS.CollateralToken {
			t.Errorf("call %d: expected CollateralToken %s, got %s",
				i, MATIC_CONTRACTS.CollateralToken.Hex(), targets[i].Hex())
		}
	}

	for i := 3; i < 7; i++ {
		if targets[i] != MATIC_CONTRACTS.ConditionalTokens {
			t.Errorf("call %d: expected ConditionalTokens %s, got %s",
				i, MATIC_CONTRACTS.ConditionalTokens.Hex(), targets[i].Hex())
		}
	}
}

func TestV2SplitPositionForEOA_UsesCtfAdapter(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2SPLIT")}
	v2 := newV2TestInstance(mock)

	_, err := v2.SplitPositionForEOA(context.Background(), testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("V2 SplitPositionForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter %s, got %s",
			MATIC_CONTRACTS.CtfCollateralAdapter.Hex(), mock.lastTo.Hex())
	}

	parsedABI, _ := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(mock.lastData, expected) {
		t.Error("V2 split calldata mismatch")
	}
}

func TestV2MergePositionsForEOA_UsesCtfAdapter(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2MERGE")}
	v2 := newV2TestInstance(mock)

	_, err := v2.MergePositionsForEOA(context.Background(), testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("V2 MergePositionsForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestV2RedeemPositionsForEOA_UsesCtfAdapter(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2REDEEM")}
	v2 := newV2TestInstance(mock)
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}

	_, err := v2.RedeemPositionsForEOA(context.Background(), testCondID, indexSets)
	if err != nil {
		t.Fatalf("V2 RedeemPositionsForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CtfCollateralAdapter {
		t.Errorf("expected target CtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestV2SplitPositionNegRiskForEOA_UsesNRAdapter(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2NRSPLIT")}
	v2 := newV2TestInstance(mock)

	_, err := v2.SplitPositionNegRiskForEOA(context.Background(), testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("V2 SplitPositionNegRiskForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.NegRiskCtfCollateralAdapter {
		t.Errorf("expected target NegRiskCtfCollateralAdapter %s, got %s",
			MATIC_CONTRACTS.NegRiskCtfCollateralAdapter.Hex(), mock.lastTo.Hex())
	}

	parsedABI, _ := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(mock.lastData, expected) {
		t.Error("V2 NegRisk split calldata mismatch")
	}
}

func TestV2MergePositionsNegRiskForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2NRMERGE")}
	v2 := newV2TestInstance(mock)

	_, err := v2.MergePositionsNegRiskForEOA(context.Background(), testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("V2 MergePositionsNegRiskForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.NegRiskCtfCollateralAdapter {
		t.Errorf("expected target NegRiskCtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestV2RedeemPositionsNegRiskForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2NRREDEEM")}
	v2 := newV2TestInstance(mock)
	indexSets := []*big.Int{big.NewInt(3), big.NewInt(4)}

	_, err := v2.RedeemPositionsNegRiskForEOA(context.Background(), testCondID, indexSets)
	if err != nil {
		t.Fatalf("V2 RedeemPositionsNegRiskForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.NegRiskCtfCollateralAdapter {
		t.Errorf("expected target NegRiskCtfCollateralAdapter, got %s", mock.lastTo.Hex())
	}
}

func TestV2WrapToPUSDForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2WRAP")}
	v2 := newV2TestInstance(mock)

	asset := MATIC_CONTRACTS.Collateral
	to := common.HexToAddress("0xRecipient")
	amount := big.NewInt(5000000)

	_, err := v2.WrapToPUSDForEOA(context.Background(), asset, to, amount)
	if err != nil {
		t.Fatalf("WrapToPUSDForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CollateralOnramp {
		t.Errorf("expected target CollateralOnramp %s, got %s",
			MATIC_CONTRACTS.CollateralOnramp.Hex(), mock.lastTo.Hex())
	}
}

func TestV2UnwrapFromPUSDForEOA(t *testing.T) {
	mock := &mockTransactionSender{retHash: common.HexToHash("0xV2UNWRAP")}
	v2 := newV2TestInstance(mock)

	asset := MATIC_CONTRACTS.Collateral
	to := common.HexToAddress("0xRecipient")
	amount := big.NewInt(3000000)

	_, err := v2.UnwrapFromPUSDForEOA(context.Background(), asset, to, amount)
	if err != nil {
		t.Fatalf("UnwrapFromPUSDForEOA: %v", err)
	}

	if mock.lastTo != MATIC_CONTRACTS.CollateralOfframp {
		t.Errorf("expected target CollateralOfframp %s, got %s",
			MATIC_CONTRACTS.CollateralOfframp.Hex(), mock.lastTo.Hex())
	}
}
