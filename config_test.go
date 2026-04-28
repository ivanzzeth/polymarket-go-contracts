package polymarketcontracts

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMATICContractsV1FieldsUnchanged(t *testing.T) {
	c := MATIC_CONTRACTS

	checks := []struct {
		name string
		got  common.Address
		want string
	}{
		{"Exchange", c.Exchange, "0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"},
		{"NegRiskAdapter", c.NegRiskAdapter, "0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"},
		{"NegRiskExchange", c.NegRiskExchange, "0xC5d563A36AE78145C45a50134d48A1215220f80a"},
		{"Collateral", c.Collateral, "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"},
		{"ConditionalTokens", c.ConditionalTokens, "0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"},
		{"SafeProxyFactory", c.SafeProxyFactory, "0xaacFeEa03eb1561C4e67d661e40682Bd20E3541b"},
	}

	for _, tc := range checks {
		if tc.got != common.HexToAddress(tc.want) {
			t.Errorf("MATIC_CONTRACTS.%s = %s, want %s", tc.name, tc.got.Hex(), tc.want)
		}
	}
}

func TestMATICContractsV2FieldsPopulated(t *testing.T) {
	c := MATIC_CONTRACTS

	checks := []struct {
		name string
		got  common.Address
		want string
	}{
		{"ExchangeV2", c.ExchangeV2, "0xE111180000d2663C0091e4f400237545B87B996B"},
		{"NegRiskExchangeV2", c.NegRiskExchangeV2, "0xe2222d279d744050d28e00520010520000310F59"},
		{"CollateralToken", c.CollateralToken, "0xC011a7E12a19f7B1f670d46F03B03f3342E82DFB"},
		{"CollateralOnramp", c.CollateralOnramp, "0x93070a847efEf7F70739046A929D47a521F5B8ee"},
		{"CollateralOfframp", c.CollateralOfframp, "0x2957922Eb93258b93368531d39fAcCA3B4dC5854"},
		{"CtfCollateralAdapter", c.CtfCollateralAdapter, "0xADa100874d00e3331D00F2007a9c336a65009718"},
		{"NegRiskCtfCollateralAdapter", c.NegRiskCtfCollateralAdapter, "0xAdA200001000ef00D07553cEE7006808F895c6F1"},
		{"PermissionedRamp", c.PermissionedRamp, "0xebC2459Ec962869ca4c0bd1E06368272732BCb08"},
	}

	for _, tc := range checks {
		if tc.got != common.HexToAddress(tc.want) {
			t.Errorf("MATIC_CONTRACTS.%s = %s, want %s", tc.name, tc.got.Hex(), tc.want)
		}
	}
}

func TestAMOYContractsV2FieldsZero(t *testing.T) {
	c := AMOY_CONTRACTS
	zero := common.Address{}

	v2Fields := []struct {
		name string
		got  common.Address
	}{
		{"ExchangeV2", c.ExchangeV2},
		{"NegRiskExchangeV2", c.NegRiskExchangeV2},
		{"CollateralToken", c.CollateralToken},
		{"CollateralOnramp", c.CollateralOnramp},
		{"CollateralOfframp", c.CollateralOfframp},
		{"CtfCollateralAdapter", c.CtfCollateralAdapter},
		{"NegRiskCtfCollateralAdapter", c.NegRiskCtfCollateralAdapter},
		{"PermissionedRamp", c.PermissionedRamp},
	}

	for _, tc := range v2Fields {
		if tc.got != zero {
			t.Errorf("AMOY_CONTRACTS.%s = %s, want zero address", tc.name, tc.got.Hex())
		}
	}
}

func TestGetContractConfig(t *testing.T) {
	matic := GetContractConfig(big.NewInt(137))
	if matic != MATIC_CONTRACTS {
		t.Error("expected MATIC_CONTRACTS for chain 137")
	}

	amoy := GetContractConfig(big.NewInt(80002))
	if amoy != AMOY_CONTRACTS {
		t.Error("expected AMOY_CONTRACTS for chain 80002")
	}
}

func TestGetContractConfigPanicsOnUnknownChain(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unknown chain ID")
		}
	}()
	GetContractConfig(big.NewInt(999))
}
