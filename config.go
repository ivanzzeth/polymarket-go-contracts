package polymarketcontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ContractConfig holds all contract addresses for Polymarket
type ContractConfig struct {
	// V1 fields
	// WARNING: "Collateral" is USDC.e (bridged USDC), NOT native USDC.
	// Polymarket V1 used USDC.e as collateral at the CTF layer.
	// This field is kept as-is for backward compatibility.
	Collateral        common.Address // USDC.e (0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174) — NOT native USDC
	ConditionalTokens common.Address
	Exchange          common.Address // V1 CTF Exchange
	NegRiskAdapter    common.Address // V1 NegRisk adapter
	NegRiskExchange   common.Address // V1 NegRisk exchange
	SafeProxyFactory  common.Address

	// V2 fields (zero-value = V2 not configured)
	USDC                        common.Address // Native USDC (0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359) — NOT USDC.e
	ExchangeV2                  common.Address // CTF Exchange V2
	NegRiskExchangeV2           common.Address // NegRisk CTF Exchange V2
	CollateralToken             common.Address // pUSD proxy — the new primary collateral
	CollateralOnramp            common.Address // USDC/USDC.e -> pUSD
	CollateralOfframp           common.Address // pUSD -> USDC/USDC.e
	CtfCollateralAdapter        common.Address // pUSD <-> CTF (regular markets)
	NegRiskCtfCollateralAdapter common.Address // pUSD <-> CTF (neg-risk markets)
	PermissionedRamp            common.Address // wrap/unwrap with signature verification
}

var AMOY_CONTRACTS = &ContractConfig{
	Exchange:          common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40"),
	NegRiskAdapter:    common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
	NegRiskExchange:   common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
	Collateral:        common.HexToAddress("0x9c4e1703476e875070ee25b56a58b008cfb8fa78"),
	ConditionalTokens: common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
	SafeProxyFactory:  common.HexToAddress("0xaacFeEa03eb1561C4e67d661e40682Bd20E3541b"),
}

var MATIC_CONTRACTS = &ContractConfig{
	// V1
	Exchange:          common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
	NegRiskAdapter:    common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
	NegRiskExchange:   common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
	Collateral:        common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"), // USDC.e
	ConditionalTokens: common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	SafeProxyFactory:  common.HexToAddress("0xaacFeEa03eb1561C4e67d661e40682Bd20E3541b"),
	// V2
	USDC:                        common.HexToAddress("0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"),
	ExchangeV2:                  common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"),
	NegRiskExchangeV2:           common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59"),
	CollateralToken:             common.HexToAddress("0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB"), // pUSD proxy
	CollateralOnramp:            common.HexToAddress("0x93070a847efEf7F70739046A929D47a521F5B8ee"),
	CollateralOfframp:           common.HexToAddress("0x2957922Eb93258b93368531d39fAcCA3B4dC5854"),
	CtfCollateralAdapter:        common.HexToAddress("0xADa100874d00e3331D00F2007a9c336a65009718"),
	NegRiskCtfCollateralAdapter: common.HexToAddress("0xAdA200001000ef00D07553cEE7006808F895c6F1"),
	PermissionedRamp:            common.HexToAddress("0xebC2459Ec962869ca4c0bd1E06368272732BCb08"),
}

func GetContractConfig(chainID *big.Int) *ContractConfig {
	if big.NewInt(137).Cmp(chainID) == 0 {
		return MATIC_CONTRACTS
	} else if big.NewInt(80002).Cmp(chainID) == 0 {
		return AMOY_CONTRACTS
	} else {
		panic("Invalid network")
	}
}
