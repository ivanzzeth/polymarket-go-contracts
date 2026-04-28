package polymarketcontracts

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	conditional_tokens "github.com/ivanzzeth/polymarket-go-contracts/contracts/conditional-tokens"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/ctf-collateral-adapter"
	"github.com/ivanzzeth/polymarket-go-contracts/contracts/erc20"
	negriskadapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-adapter"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-ctf-collateral-adapter"
)

var (
	testCTF        = common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045")
	testCollateral = common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
	testAdapter    = common.HexToAddress("0xADa100874d00e3331D00F2007a9c336a65009718")
	testNRAdapter  = common.HexToAddress("0xAdA200001000ef00D07553cEE7006808F895c6F1")
	testV1Adapter  = common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296")
	testCondID     = [32]byte{0x01, 0x02, 0x03}
	testPartition  = []*big.Int{big.NewInt(1), big.NewInt(2)}
	testAmount     = big.NewInt(1000000)
)

func TestBuildSplitPositionCall(t *testing.T) {
	call, err := buildSplitPositionCall(testCTF, testCollateral, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != testCTF {
		t.Errorf("expected target %s, got %s", testCTF.Hex(), call.Target.Hex())
	}
	if call.Value.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("expected zero value, got %s", call.Value)
	}

	parsedABI, _ := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", testCollateral, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Errorf("calldata mismatch\nexpected: %x\ngot:      %x", expected, call.Calldata)
	}
}

func TestBuildMergePositionsCall(t *testing.T) {
	call, err := buildMergePositionsCall(testCTF, testCollateral, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != testCTF {
		t.Errorf("expected target %s", testCTF.Hex())
	}

	parsedABI, _ := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	expected, _ := parsedABI.Pack("mergePositions", testCollateral, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for mergePositions")
	}
}

func TestBuildRedeemPositionsCall(t *testing.T) {
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}
	call, err := buildRedeemPositionsCall(testCTF, testCollateral, testCondID, indexSets)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	expected, _ := parsedABI.Pack("redeemPositions", testCollateral, [32]byte{}, testCondID, indexSets)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for redeemPositions")
	}
}

func TestBuildSplitNegRiskCall(t *testing.T) {
	call, err := buildSplitNegRiskCall(testV1Adapter, testCondID, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != testV1Adapter {
		t.Errorf("expected target %s", testV1Adapter.Hex())
	}

	parsedABI, _ := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition0", testCondID, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for NegRisk splitPosition0")
	}
}

func TestBuildMergeNegRiskCall(t *testing.T) {
	call, err := buildMergeNegRiskCall(testV1Adapter, testCondID, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("mergePositions0", testCondID, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for NegRisk mergePositions0")
	}
}

func TestBuildRedeemNegRiskCall(t *testing.T) {
	amounts := []*big.Int{big.NewInt(500), big.NewInt(600)}
	call, err := buildRedeemNegRiskCall(testV1Adapter, testCondID, amounts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("redeemPositions", testCondID, amounts)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for NegRisk redeemPositions")
	}
}

func TestBuildAdapterSplitCall(t *testing.T) {
	call, err := buildAdapterSplitCall(testAdapter, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != testAdapter {
		t.Errorf("expected target %s", testAdapter.Hex())
	}

	parsedABI, _ := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for adapter splitPosition")
	}
}

func TestBuildAdapterMergeCall(t *testing.T) {
	call, err := buildAdapterMergeCall(testAdapter, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("mergePositions", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for adapter mergePositions")
	}
}

func TestBuildAdapterRedeemCall(t *testing.T) {
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}
	call, err := buildAdapterRedeemCall(testAdapter, testCondID, indexSets)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("redeemPositions", common.Address{}, [32]byte{}, testCondID, indexSets)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for adapter redeemPositions")
	}
}

func TestBuildNegRiskAdapterSplitCall(t *testing.T) {
	call, err := buildNegRiskAdapterSplitCall(testNRAdapter, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != testNRAdapter {
		t.Errorf("expected target %s", testNRAdapter.Hex())
	}

	parsedABI, _ := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for neg-risk adapter splitPosition")
	}
}

func TestBuildNegRiskAdapterMergeCall(t *testing.T) {
	call, err := buildNegRiskAdapterMergeCall(testNRAdapter, testCondID, testPartition, testAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("mergePositions", common.Address{}, [32]byte{}, testCondID, testPartition, testAmount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for neg-risk adapter mergePositions")
	}
}

func TestBuildNegRiskAdapterRedeemCall(t *testing.T) {
	indexSets := []*big.Int{big.NewInt(3), big.NewInt(4)}
	call, err := buildNegRiskAdapterRedeemCall(testNRAdapter, testCondID, indexSets)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parsedABI, _ := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	expected, _ := parsedABI.Pack("redeemPositions", common.Address{}, [32]byte{}, testCondID, indexSets)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for neg-risk adapter redeemPositions")
	}
}

func TestBuildERC20ApproveCall(t *testing.T) {
	token := common.HexToAddress("0xToken")
	spender := common.HexToAddress("0xSpender")
	amount := big.NewInt(999)

	call, err := buildERC20ApproveCall(token, spender, amount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != token {
		t.Errorf("expected target %s", token.Hex())
	}

	parsedABI, _ := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	expected, _ := parsedABI.Pack("approve", spender, amount)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for ERC20 approve")
	}
}

func TestBuildSetApprovalForAllCall(t *testing.T) {
	ctf := testCTF
	operator := common.HexToAddress("0xOp")

	call, err := buildSetApprovalForAllCall(ctf, operator, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != ctf {
		t.Errorf("expected target %s", ctf.Hex())
	}

	parsedABI, _ := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	expected, _ := parsedABI.Pack("setApprovalForAll", operator, true)
	if !bytes.Equal(call.Calldata, expected) {
		t.Error("calldata mismatch for setApprovalForAll")
	}
}

func TestBuildWrapCall(t *testing.T) {
	onramp := common.HexToAddress("0xOnramp")
	asset := common.HexToAddress("0xAsset")
	to := common.HexToAddress("0xTo")
	amount := big.NewInt(5000000)

	call, err := buildWrapCall(onramp, asset, to, amount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != onramp {
		t.Errorf("expected target %s", onramp.Hex())
	}
	if call.Value.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("expected zero value")
	}
}

func TestBuildUnwrapCall(t *testing.T) {
	offramp := common.HexToAddress("0xOfframp")
	asset := common.HexToAddress("0xAsset")
	to := common.HexToAddress("0xTo")
	amount := big.NewInt(3000000)

	call, err := buildUnwrapCall(offramp, asset, to, amount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if call.Target != offramp {
		t.Errorf("expected target %s", offramp.Hex())
	}
}

func TestAllCallBuildersReturnZeroValue(t *testing.T) {
	builders := []struct {
		name string
		fn   func() (contractCall, error)
	}{
		{"splitPosition", func() (contractCall, error) {
			return buildSplitPositionCall(testCTF, testCollateral, testCondID, testPartition, testAmount)
		}},
		{"mergePositions", func() (contractCall, error) {
			return buildMergePositionsCall(testCTF, testCollateral, testCondID, testPartition, testAmount)
		}},
		{"redeemPositions", func() (contractCall, error) {
			return buildRedeemPositionsCall(testCTF, testCollateral, testCondID, testPartition)
		}},
		{"splitNegRisk", func() (contractCall, error) {
			return buildSplitNegRiskCall(testV1Adapter, testCondID, testAmount)
		}},
		{"adapterSplit", func() (contractCall, error) {
			return buildAdapterSplitCall(testAdapter, testCondID, testPartition, testAmount)
		}},
		{"negRiskAdapterSplit", func() (contractCall, error) {
			return buildNegRiskAdapterSplitCall(testNRAdapter, testCondID, testPartition, testAmount)
		}},
		{"wrap", func() (contractCall, error) {
			return buildWrapCall(testAdapter, testCollateral, testCTF, testAmount)
		}},
		{"unwrap", func() (contractCall, error) {
			return buildUnwrapCall(testAdapter, testCollateral, testCTF, testAmount)
		}},
		{"erc20Approve", func() (contractCall, error) {
			return buildERC20ApproveCall(testCollateral, testAdapter, testAmount)
		}},
		{"setApprovalForAll", func() (contractCall, error) {
			return buildSetApprovalForAllCall(testCTF, testAdapter, true)
		}},
	}

	for _, b := range builders {
		t.Run(b.name, func(t *testing.T) {
			call, err := b.fn()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if call.Value == nil || call.Value.Cmp(big.NewInt(0)) != 0 {
				t.Errorf("expected zero value, got %v", call.Value)
			}
		})
	}
}
