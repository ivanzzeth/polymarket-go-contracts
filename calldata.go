package polymarketcontracts

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	collateral_offramp "github.com/ivanzzeth/polymarket-go-contracts/contracts/collateral-offramp"
	collateral_onramp "github.com/ivanzzeth/polymarket-go-contracts/contracts/collateral-onramp"
	conditional_tokens "github.com/ivanzzeth/polymarket-go-contracts/contracts/conditional-tokens"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/ctf-collateral-adapter"
	"github.com/ivanzzeth/polymarket-go-contracts/contracts/erc20"
	negriskadapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-adapter"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-ctf-collateral-adapter"
)

// V1 calldata builders — regular markets (CTF direct)

func buildSplitPositionCall(ctf, collateral common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{}
	calldata, err := parsedABI.Pack("splitPosition", collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack splitPosition calldata: %w", err)
	}
	return contractCall{Target: ctf, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildMergePositionsCall(ctf, collateral common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{}
	calldata, err := parsedABI.Pack("mergePositions", collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack mergePositions calldata: %w", err)
	}
	return contractCall{Target: ctf, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildRedeemPositionsCall(ctf, collateral common.Address, conditionId [32]byte, indexSets []*big.Int) (contractCall, error) {
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{}
	calldata, err := parsedABI.Pack("redeemPositions", collateral, parentCollectionId, conditionId, indexSets)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}
	return contractCall{Target: ctf, Calldata: calldata, Value: big.NewInt(0)}, nil
}

// V1 calldata builders — NegRisk markets (NegRiskAdapter)

func buildSplitNegRiskCall(adapter common.Address, conditionId [32]byte, amount *big.Int) (contractCall, error) {
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition0", conditionId, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack splitPosition0 calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildMergeNegRiskCall(adapter common.Address, conditionId [32]byte, amount *big.Int) (contractCall, error) {
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions0", conditionId, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack mergePositions0 calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildRedeemNegRiskCall(adapter common.Address, conditionId [32]byte, amounts []*big.Int) (contractCall, error) {
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", conditionId, amounts)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

// V2 calldata builders — pUSD via CtfCollateralAdapter (regular markets)

func buildAdapterSplitCall(adapter common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse CtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack adapter splitPosition calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildAdapterMergeCall(adapter common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse CtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions", common.Address{}, [32]byte{}, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack adapter mergePositions calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildAdapterRedeemCall(adapter common.Address, conditionId [32]byte, indexSets []*big.Int) (contractCall, error) {
	parsedABI, err := ctf_collateral_adapter.CtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse CtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", common.Address{}, [32]byte{}, conditionId, indexSets)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack adapter redeemPositions calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

// V2 calldata builders — pUSD via NegRiskCtfCollateralAdapter (neg-risk markets)

func buildNegRiskAdapterSplitCall(adapter common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskCtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition", common.Address{}, [32]byte{}, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack neg-risk adapter splitPosition calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildNegRiskAdapterMergeCall(adapter common.Address, conditionId [32]byte, partition []*big.Int, amount *big.Int) (contractCall, error) {
	parsedABI, err := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskCtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions", common.Address{}, [32]byte{}, conditionId, partition, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack neg-risk adapter mergePositions calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildNegRiskAdapterRedeemCall(adapter common.Address, conditionId [32]byte, indexSets []*big.Int) (contractCall, error) {
	parsedABI, err := neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse NegRiskCtfCollateralAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", common.Address{}, [32]byte{}, conditionId, indexSets)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack neg-risk adapter redeemPositions calldata: %w", err)
	}
	return contractCall{Target: adapter, Calldata: calldata, Value: big.NewInt(0)}, nil
}

// Wrap/Unwrap calldata builders

func buildWrapCall(onramp common.Address, asset, to common.Address, amount *big.Int) (contractCall, error) {
	parsedABI, err := collateral_onramp.CollateralOnrampMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse CollateralOnramp ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("wrap", asset, to, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack wrap calldata: %w", err)
	}
	return contractCall{Target: onramp, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildUnwrapCall(offramp common.Address, asset, to common.Address, amount *big.Int) (contractCall, error) {
	parsedABI, err := collateral_offramp.CollateralOfframpMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse CollateralOfframp ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("unwrap", asset, to, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack unwrap calldata: %w", err)
	}
	return contractCall{Target: offramp, Calldata: calldata, Value: big.NewInt(0)}, nil
}

// ERC-20 / ERC-1155 approval calldata builders

func buildERC20ApproveCall(token, spender common.Address, amount *big.Int) (contractCall, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("approve", spender, amount)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack approve calldata: %w", err)
	}
	return contractCall{Target: token, Calldata: calldata, Value: big.NewInt(0)}, nil
}

func buildSetApprovalForAllCall(ctf, operator common.Address, approved bool) (contractCall, error) {
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		return contractCall{}, fmt.Errorf("failed to pack setApprovalForAll calldata: %w", err)
	}
	return contractCall{Target: ctf, Calldata: calldata, Value: big.NewInt(0)}, nil
}
