package polymarketcontracts

import (
	"bytes"
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/v2/contracts/ctf-collateral-adapter"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/v2/contracts/neg-risk-ctf-collateral-adapter"
)

func newV2TestInstance(mock *mockTransactionSender) *ContractInterfaceV2 {
	v := &ContractInterfaceV2{
		config: MATIC_CONTRACTS,
		executor: &txExecutor{
			txSender: mock,
		},
		tokenStatus: map[common.Address]*TokenStatus{},
		statusTTL:   5 * time.Minute,
	}
	now := time.Now()
	v.tokenStatus[MATIC_CONTRACTS.Collateral] = &TokenStatus{WrapEnabled: true, UnwrapEnabled: true, LastChecked: now}
	v.tokenStatus[MATIC_CONTRACTS.USDC] = &TokenStatus{WrapEnabled: true, UnwrapEnabled: true, LastChecked: now}
	return v
}

func TestV2BuildEnableTradingCalls_CallCount(t *testing.T) {
	v2 := &ContractInterfaceV2{config: MATIC_CONTRACTS}

	// Zero-value info: all allowances nil/zero, all approvals false → all calls needed
	info := &V2BalanceInfo{}
	calls, err := v2.buildEnableTradingCalls(info)
	if err != nil {
		t.Fatalf("buildEnableTradingCalls: %v", err)
	}

	// 1 USDC.e approve + 1 USDC approve + 6 pUSD approves + 4 CTF setApprovalForAll = 12
	if len(calls) != 12 {
		t.Fatalf("expected 12 calls with zero allowances, got %d", len(calls))
	}
}

func TestV2BuildEnableTradingCalls_Targets(t *testing.T) {
	v2 := &ContractInterfaceV2{config: MATIC_CONTRACTS}

	info := &V2BalanceInfo{}
	calls, err := v2.buildEnableTradingCalls(info)
	if err != nil {
		t.Fatalf("buildEnableTradingCalls: %v", err)
	}

	if len(calls) != 12 {
		t.Fatalf("expected 12 calls, got %d", len(calls))
	}

	// calls[0]: USDC.e approve → target is Collateral (USDC.e token)
	if calls[0].Target != MATIC_CONTRACTS.Collateral {
		t.Errorf("call 0: expected Collateral %s, got %s",
			MATIC_CONTRACTS.Collateral.Hex(), calls[0].Target.Hex())
	}
	// calls[1]: USDC approve → target is USDC token
	if calls[1].Target != MATIC_CONTRACTS.USDC {
		t.Errorf("call 1: expected USDC %s, got %s",
			MATIC_CONTRACTS.USDC.Hex(), calls[1].Target.Hex())
	}
	// calls[2..7]: pUSD approves → target is CollateralToken (pUSD)
	for i := 2; i < 8; i++ {
		if calls[i].Target != MATIC_CONTRACTS.CollateralToken {
			t.Errorf("call %d: expected CollateralToken %s, got %s",
				i, MATIC_CONTRACTS.CollateralToken.Hex(), calls[i].Target.Hex())
		}
	}
	// calls[8..11]: CTF setApprovalForAll → target is ConditionalTokens
	for i := 8; i < 12; i++ {
		if calls[i].Target != MATIC_CONTRACTS.ConditionalTokens {
			t.Errorf("call %d: expected ConditionalTokens %s, got %s",
				i, MATIC_CONTRACTS.ConditionalTokens.Hex(), calls[i].Target.Hex())
		}
	}
}

func TestV2BuildEnableTradingCalls_SkipsApproved(t *testing.T) {
	v2 := &ContractInterfaceV2{config: MATIC_CONTRACTS}

	maxApproval := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	info := &V2BalanceInfo{
		USDCEAllowanceOnramp:           maxApproval,
		USDCAllowanceOnramp:            maxApproval,
		PUSDAllowanceExchangeV2:        maxApproval,
		PUSDAllowanceNegRiskExchangeV2: maxApproval,
		PUSDAllowanceNegRiskAdapter:    maxApproval,
		PUSDAllowanceCtfAdapter:        maxApproval,
		PUSDAllowanceNegRiskCtfAdapter: maxApproval,
		PUSDAllowanceOfframp:           maxApproval,
		CTFApprovedExchangeV2:                  true,
		CTFApprovedNegRiskExchangeV2:           true,
		CTFApprovedCtfCollateralAdapter:        true,
		CTFApprovedNegRiskCtfCollateralAdapter: true,
	}
	calls, err := v2.buildEnableTradingCalls(info)
	if err != nil {
		t.Fatalf("buildEnableTradingCalls: %v", err)
	}

	if len(calls) != 0 {
		t.Fatalf("expected 0 calls when all approved, got %d", len(calls))
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
