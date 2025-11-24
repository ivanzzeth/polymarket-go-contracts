package polymarketcontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ContractConfig holds all contract addresses for Polymarket
type ContractConfig struct {
	Collateral        common.Address // USDC token address
	ConditionalTokens common.Address // CTF contract address
	Exchange          common.Address // Exchange contract address
	NegRiskAdapter    common.Address // NegRisk adapter address
	NegRiskExchange   common.Address // NegRisk exchange address
	SafeProxyFactory  common.Address // Safe Proxy Factory address
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
	Exchange:          common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
	NegRiskAdapter:    common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
	NegRiskExchange:   common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
	Collateral:        common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
	ConditionalTokens: common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	SafeProxyFactory:  common.HexToAddress("0xaacFeEa03eb1561C4e67d661e40682Bd20E3541b"),
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
