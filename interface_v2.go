package polymarketcontracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethclient"
	collateral_offramp "github.com/ivanzzeth/polymarket-go-contracts/contracts/collateral-offramp"
	collateral_onramp "github.com/ivanzzeth/polymarket-go-contracts/contracts/collateral-onramp"
	collateral_token "github.com/ivanzzeth/polymarket-go-contracts/contracts/collateral-token"
	conditional_tokens "github.com/ivanzzeth/polymarket-go-contracts/contracts/conditional-tokens"
	ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/ctf-collateral-adapter"
	"github.com/ivanzzeth/polymarket-go-contracts/contracts/erc20"
	exchange_v2 "github.com/ivanzzeth/polymarket-go-contracts/contracts/exchange-v2"
	neg_risk_ctf_collateral_adapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-ctf-collateral-adapter"
	neg_risk_v2 "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-v2"
	permissioned_ramp "github.com/ivanzzeth/polymarket-go-contracts/contracts/permissioned-ramp"
	"github.com/ivanzzeth/polymarket-go-contracts/sender"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

// V2BalanceInfo holds pUSD and USDC.e balances plus V2-relevant allowances and approvals.
type V2BalanceInfo struct {
	PUSDBalance  *big.Int
	USDCBalance  *big.Int
	USDCEBalance *big.Int

	USDCAllowanceOnramp  *big.Int
	USDCEAllowanceOnramp *big.Int

	PUSDAllowanceCtfAdapter        *big.Int
	PUSDAllowanceNegRiskCtfAdapter *big.Int
	PUSDAllowanceOfframp           *big.Int

	CTFApprovedExchangeV2                  bool
	CTFApprovedNegRiskExchangeV2           bool
	CTFApprovedCtfCollateralAdapter        bool
	CTFApprovedNegRiskCtfCollateralAdapter bool
}

// ContractInterfaceV2 provides a clean V2 API where pUSD is the default collateral.
// It delegates to the same shared txExecutor and calldata builders used by V1.
type ContractInterfaceV2 struct {
	config   *ContractConfig
	client   ethclient.EthClientInterface
	executor *txExecutor

	exchangeV2                  *exchange_v2.ExchangeV2
	negRiskExchangeV2           *neg_risk_v2.NegRiskV2
	collateralToken             *collateral_token.CollateralToken
	usdc                        *erc20.Erc20
	usdce                       *erc20.Erc20
	conditionalTokens           *conditional_tokens.ConditionalTokens
	collateralOnramp            *collateral_onramp.CollateralOnramp
	collateralOfframp           *collateral_offramp.CollateralOfframp
	ctfCollateralAdapter        *ctf_collateral_adapter.CtfCollateralAdapter
	negRiskCtfCollateralAdapter *neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapter
	permissionedRamp            *permissioned_ramp.PermissionedRamp
}

// NewContractInterfaceV2 creates a V2 interface. All V2 contract addresses in config must be non-zero.
func NewContractInterfaceV2(
	client ethclient.EthClientInterface,
	config *ContractConfig,
	txSender sender.TransactionSender,
	getSafeAddr func(eoa common.Address) (common.Address, error),
	execSafeTx func(safeSigner signer.SafeTradingSigner, chainID *big.Int, safeAddr, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int) (common.Hash, error),
) (*ContractInterfaceV2, error) {
	if config.ExchangeV2 == (common.Address{}) {
		return nil, fmt.Errorf("V2 not configured: ExchangeV2 address is zero")
	}

	exchV2, err := exchange_v2.NewExchangeV2(config.ExchangeV2, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExchangeV2 binding: %w", err)
	}
	nrV2, err := neg_risk_v2.NewNegRiskV2(config.NegRiskExchangeV2, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create NegRiskV2 binding: %w", err)
	}
	ct, err := collateral_token.NewCollateralToken(config.CollateralToken, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CollateralToken binding: %w", err)
	}
	usdce, err := erc20.NewErc20(config.Collateral, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create USDC.e binding: %w", err)
	}
	var usdc *erc20.Erc20
	if config.USDC != (common.Address{}) {
		usdc, err = erc20.NewErc20(config.USDC, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create USDC binding: %w", err)
		}
	}
	ctf, err := conditional_tokens.NewConditionalTokens(config.ConditionalTokens, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create ConditionalTokens binding: %w", err)
	}
	onramp, err := collateral_onramp.NewCollateralOnramp(config.CollateralOnramp, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CollateralOnramp binding: %w", err)
	}
	offramp, err := collateral_offramp.NewCollateralOfframp(config.CollateralOfframp, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CollateralOfframp binding: %w", err)
	}
	ctfAdapter, err := ctf_collateral_adapter.NewCtfCollateralAdapter(config.CtfCollateralAdapter, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CtfCollateralAdapter binding: %w", err)
	}
	nrCtfAdapter, err := neg_risk_ctf_collateral_adapter.NewNegRiskCtfCollateralAdapter(config.NegRiskCtfCollateralAdapter, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create NegRiskCtfCollateralAdapter binding: %w", err)
	}
	pr, err := permissioned_ramp.NewPermissionedRamp(config.PermissionedRamp, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create PermissionedRamp binding: %w", err)
	}

	return &ContractInterfaceV2{
		config: config,
		client: client,
		executor: &txExecutor{
			client:      client,
			txSender:    txSender,
			getSafeAddr: getSafeAddr,
			execSafeTx:  execSafeTx,
		},
		exchangeV2:                  exchV2,
		negRiskExchangeV2:           nrV2,
		collateralToken:             ct,
		usdc:                        usdc,
		usdce:                       usdce,
		conditionalTokens:           ctf,
		collateralOnramp:            onramp,
		collateralOfframp:           offramp,
		ctfCollateralAdapter:        ctfAdapter,
		negRiskCtfCollateralAdapter: nrCtfAdapter,
		permissionedRamp:            pr,
	}, nil
}

// --- Read operations ---

// GetBalances returns pUSD/USDC.e balances and V2-relevant allowances/approvals.
func (v *ContractInterfaceV2) GetBalances(ctx context.Context, address common.Address) (*V2BalanceInfo, error) {
	return v.GetBalancesAtBlock(ctx, address, nil)
}

// GetBalancesAtBlock returns balances at a specific block (nil = latest).
func (v *ContractInterfaceV2) GetBalancesAtBlock(ctx context.Context, address common.Address, blockNumber *big.Int) (*V2BalanceInfo, error) {
	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}
	info := &V2BalanceInfo{}

	pusdBal, err := v.collateralToken.BalanceOf(opts, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get pUSD balance: %w", err)
	}
	info.PUSDBalance = pusdBal

	usdceBal, err := v.usdce.BalanceOf(opts, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC.e balance: %w", err)
	}
	info.USDCEBalance = usdceBal

	info.USDCEAllowanceOnramp, err = v.usdce.Allowance(opts, address, v.config.CollateralOnramp)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC.e allowance for CollateralOnramp: %w", err)
	}

	if v.usdc != nil {
		usdcBal, err := v.usdc.BalanceOf(opts, address)
		if err != nil {
			return nil, fmt.Errorf("failed to get USDC balance: %w", err)
		}
		info.USDCBalance = usdcBal

		info.USDCAllowanceOnramp, err = v.usdc.Allowance(opts, address, v.config.CollateralOnramp)
		if err != nil {
			return nil, fmt.Errorf("failed to get USDC allowance for CollateralOnramp: %w", err)
		}
	}

	info.PUSDAllowanceCtfAdapter, err = v.collateralToken.Allowance(opts, address, v.config.CtfCollateralAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to get pUSD allowance for CtfCollateralAdapter: %w", err)
	}
	info.PUSDAllowanceNegRiskCtfAdapter, err = v.collateralToken.Allowance(opts, address, v.config.NegRiskCtfCollateralAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to get pUSD allowance for NegRiskCtfCollateralAdapter: %w", err)
	}
	info.PUSDAllowanceOfframp, err = v.collateralToken.Allowance(opts, address, v.config.CollateralOfframp)
	if err != nil {
		return nil, fmt.Errorf("failed to get pUSD allowance for CollateralOfframp: %w", err)
	}

	info.CTFApprovedExchangeV2, err = v.conditionalTokens.IsApprovedForAll(opts, address, v.config.ExchangeV2)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for ExchangeV2: %w", err)
	}
	info.CTFApprovedNegRiskExchangeV2, err = v.conditionalTokens.IsApprovedForAll(opts, address, v.config.NegRiskExchangeV2)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskExchangeV2: %w", err)
	}
	info.CTFApprovedCtfCollateralAdapter, err = v.conditionalTokens.IsApprovedForAll(opts, address, v.config.CtfCollateralAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for CtfCollateralAdapter: %w", err)
	}
	info.CTFApprovedNegRiskCtfCollateralAdapter, err = v.conditionalTokens.IsApprovedForAll(opts, address, v.config.NegRiskCtfCollateralAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskCtfCollateralAdapter: %w", err)
	}

	return info, nil
}

// --- Getters for V2 contract bindings ---

func (v *ContractInterfaceV2) ExchangeV2() *exchange_v2.ExchangeV2                   { return v.exchangeV2 }
func (v *ContractInterfaceV2) NegRiskExchangeV2() *neg_risk_v2.NegRiskV2             { return v.negRiskExchangeV2 }
func (v *ContractInterfaceV2) CollateralToken() *collateral_token.CollateralToken     { return v.collateralToken }
func (v *ContractInterfaceV2) NativeUSDC() *erc20.Erc20                               { return v.usdc }
func (v *ContractInterfaceV2) USDCE() *erc20.Erc20                                   { return v.usdce }
func (v *ContractInterfaceV2) ConditionalTokens() *conditional_tokens.ConditionalTokens { return v.conditionalTokens }
func (v *ContractInterfaceV2) CollateralOnramp() *collateral_onramp.CollateralOnramp  { return v.collateralOnramp }
func (v *ContractInterfaceV2) CollateralOfframp() *collateral_offramp.CollateralOfframp { return v.collateralOfframp }
func (v *ContractInterfaceV2) CtfCollateralAdapter() *ctf_collateral_adapter.CtfCollateralAdapter { return v.ctfCollateralAdapter }
func (v *ContractInterfaceV2) NegRiskCtfCollateralAdapter() *neg_risk_ctf_collateral_adapter.NegRiskCtfCollateralAdapter { return v.negRiskCtfCollateralAdapter }
func (v *ContractInterfaceV2) PermissionedRamp() *permissioned_ramp.PermissionedRamp { return v.permissionedRamp }

// --- EnableTrading ---

// EnableTradingForEOA approves pUSD to adapters/offramp and sets CTF approvals for V2 exchanges/adapters.
func (v *ContractInterfaceV2) EnableTradingForEOA(ctx context.Context) ([]common.Hash, error) {
	calls, err := v.enableTradingCalls()
	if err != nil {
		return nil, err
	}
	return v.executor.executeBatchEOA(calls)
}

// EnableTradingForSafe approves pUSD to adapters/offramp and sets CTF approvals for V2 exchanges/adapters via Safe.
func (v *ContractInterfaceV2) EnableTradingForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int) ([]common.Hash, error) {
	calls, err := v.enableTradingCalls()
	if err != nil {
		return nil, err
	}
	return v.executor.executeBatchSafe(safeSigner, chainID, calls)
}

func (v *ContractInterfaceV2) enableTradingCalls() ([]contractCall, error) {
	maxApproval := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

	var calls []contractCall

	// USDC.e approve → CollateralOnramp (needed for wrap: USDC.e → pUSD)
	usdceApproveCall, err := buildERC20ApproveCall(v.config.Collateral, v.config.CollateralOnramp, maxApproval)
	if err != nil {
		return nil, fmt.Errorf("failed to build USDC.e approve for CollateralOnramp: %w", err)
	}
	calls = append(calls, usdceApproveCall)

	// USDC (native) approve → CollateralOnramp (needed for wrap: USDC → pUSD)
	if v.config.USDC != (common.Address{}) {
		usdcApproveCall, err := buildERC20ApproveCall(v.config.USDC, v.config.CollateralOnramp, maxApproval)
		if err != nil {
			return nil, fmt.Errorf("failed to build USDC approve for CollateralOnramp: %w", err)
		}
		calls = append(calls, usdcApproveCall)
	}

	// pUSD approve → adapters + offramp
	pUSDApproveTargets := []common.Address{
		v.config.CtfCollateralAdapter,
		v.config.NegRiskCtfCollateralAdapter,
		v.config.CollateralOfframp,
	}
	for _, spender := range pUSDApproveTargets {
		call, err := buildERC20ApproveCall(v.config.CollateralToken, spender, maxApproval)
		if err != nil {
			return nil, fmt.Errorf("failed to build pUSD approve for %s: %w", spender.Hex(), err)
		}
		calls = append(calls, call)
	}

	// CTF setApprovalForAll → exchanges + adapters
	ctfApprovalTargets := []common.Address{
		v.config.ExchangeV2,
		v.config.NegRiskExchangeV2,
		v.config.CtfCollateralAdapter,
		v.config.NegRiskCtfCollateralAdapter,
	}
	for _, operator := range ctfApprovalTargets {
		call, err := buildSetApprovalForAllCall(v.config.ConditionalTokens, operator, true)
		if err != nil {
			return nil, fmt.Errorf("failed to build CTF setApprovalForAll for %s: %w", operator.Hex(), err)
		}
		calls = append(calls, call)
	}
	return calls, nil
}

// --- Wrap / Unwrap ---

// validateWrapAsset checks if the asset is a supported collateral token (USDC or USDC.e).
func (v *ContractInterfaceV2) validateWrapAsset(asset common.Address) error {
	if asset != v.config.Collateral && asset != v.config.USDC {
		return fmt.Errorf("unsupported wrap asset %s: only USDC (%s) and USDC.e (%s) are supported",
			asset.Hex(), v.config.USDC.Hex(), v.config.Collateral.Hex())
	}
	return nil
}

// WrapToPUSDForEOA wraps an asset (USDC or USDC.e) to pUSD via the CollateralOnramp.
func (v *ContractInterfaceV2) WrapToPUSDForEOA(ctx context.Context, asset common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	if err := v.validateWrapAsset(asset); err != nil {
		return common.Hash{}, err
	}
	call, err := buildWrapCall(v.config.CollateralOnramp, asset, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// WrapToPUSDForSafe wraps an asset to pUSD via Safe.
func (v *ContractInterfaceV2) WrapToPUSDForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, asset common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	if err := v.validateWrapAsset(asset); err != nil {
		return common.Hash{}, err
	}
	call, err := buildWrapCall(v.config.CollateralOnramp, asset, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// UnwrapFromPUSDForEOA unwraps pUSD to an asset (USDC or USDC.e) via the CollateralOfframp.
func (v *ContractInterfaceV2) UnwrapFromPUSDForEOA(ctx context.Context, asset common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	if err := v.validateWrapAsset(asset); err != nil {
		return common.Hash{}, err
	}
	call, err := buildUnwrapCall(v.config.CollateralOfframp, asset, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// UnwrapFromPUSDForSafe unwraps pUSD to an asset via Safe.
func (v *ContractInterfaceV2) UnwrapFromPUSDForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, asset common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	if err := v.validateWrapAsset(asset); err != nil {
		return common.Hash{}, err
	}
	call, err := buildUnwrapCall(v.config.CollateralOfframp, asset, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// --- Split / Merge / Redeem (regular markets via CtfCollateralAdapter) ---

// SplitPositionForEOA splits pUSD into conditional tokens via the CtfCollateralAdapter.
func (v *ContractInterfaceV2) SplitPositionForEOA(ctx context.Context, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildAdapterSplitCall(v.config.CtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// SplitPositionForSafe splits pUSD into conditional tokens via Safe.
func (v *ContractInterfaceV2) SplitPositionForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildAdapterSplitCall(v.config.CtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// MergePositionsForEOA merges conditional tokens back into pUSD via the CtfCollateralAdapter.
func (v *ContractInterfaceV2) MergePositionsForEOA(ctx context.Context, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildAdapterMergeCall(v.config.CtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// MergePositionsForSafe merges conditional tokens back into pUSD via Safe.
func (v *ContractInterfaceV2) MergePositionsForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildAdapterMergeCall(v.config.CtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// RedeemPositionsForEOA redeems conditional tokens for pUSD via the CtfCollateralAdapter.
func (v *ContractInterfaceV2) RedeemPositionsForEOA(ctx context.Context, conditionId [32]byte, indexSets []*big.Int) (common.Hash, error) {
	call, err := buildAdapterRedeemCall(v.config.CtfCollateralAdapter, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// RedeemPositionsForSafe redeems conditional tokens for pUSD via Safe.
func (v *ContractInterfaceV2) RedeemPositionsForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, indexSets []*big.Int) (common.Hash, error) {
	call, err := buildAdapterRedeemCall(v.config.CtfCollateralAdapter, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// --- Split / Merge / Redeem (neg-risk markets via NegRiskCtfCollateralAdapter) ---

// SplitPositionNegRiskForEOA splits pUSD into neg-risk conditional tokens via the NegRiskCtfCollateralAdapter.
func (v *ContractInterfaceV2) SplitPositionNegRiskForEOA(ctx context.Context, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterSplitCall(v.config.NegRiskCtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// SplitPositionNegRiskForSafe splits pUSD into neg-risk conditional tokens via Safe.
func (v *ContractInterfaceV2) SplitPositionNegRiskForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterSplitCall(v.config.NegRiskCtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// MergePositionsNegRiskForEOA merges neg-risk conditional tokens back into pUSD.
func (v *ContractInterfaceV2) MergePositionsNegRiskForEOA(ctx context.Context, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterMergeCall(v.config.NegRiskCtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// MergePositionsNegRiskForSafe merges neg-risk conditional tokens back into pUSD via Safe.
func (v *ContractInterfaceV2) MergePositionsNegRiskForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, partition []*big.Int, amount *big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterMergeCall(v.config.NegRiskCtfCollateralAdapter, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}

// RedeemPositionsNegRiskForEOA redeems neg-risk conditional tokens for pUSD.
func (v *ContractInterfaceV2) RedeemPositionsNegRiskForEOA(ctx context.Context, conditionId [32]byte, indexSets []*big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterRedeemCall(v.config.NegRiskCtfCollateralAdapter, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeEOA(call)
}

// RedeemPositionsNegRiskForSafe redeems neg-risk conditional tokens for pUSD via Safe.
func (v *ContractInterfaceV2) RedeemPositionsNegRiskForSafe(ctx context.Context, safeSigner signer.SafeTradingSigner, chainID *big.Int, conditionId [32]byte, indexSets []*big.Int) (common.Hash, error) {
	call, err := buildNegRiskAdapterRedeemCall(v.config.NegRiskCtfCollateralAdapter, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, err
	}
	return v.executor.executeSafe(safeSigner, chainID, call)
}
