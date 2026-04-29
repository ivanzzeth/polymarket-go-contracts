package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

func main() {
	polygonChainID := big.NewInt(137)

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to dial ethclient: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	eoaAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("EOA Address: %s", eoaAddress.Hex())

	safeSigner, err := signer.NewSafeTradingPrivateKeySigner(polygonChainID, client, privateKey)
	if err != nil {
		log.Fatalf("Failed to create safe signer: %v", err)
	}

	config := polymarketcontracts.GetContractConfig(polygonChainID)

	// V2 is self-contained, no V1 dependency needed
	v2, err := polymarketcontracts.NewContractInterfaceV2(
		client,
		config,
		safeSigner,
		polygonChainID,
	)
	if err != nil {
		log.Fatalf("Failed to create V2 contract interface: %v", err)
	}

	// Get Safe address (V2 has its own GetSafeAddress method)
	safeAddr, err := v2.GetSafeAddress(eoaAddress)
	if err != nil {
		log.Fatalf("Failed to create V2 contract interface: %v", err)
	}

	info, err := v2.GetBalances(context.Background(), safeAddr)
	if err != nil {
		log.Fatalf("Failed to get V2 balances: %v", err)
	}

	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	fmtBal := func(raw *big.Int) string {
		whole := new(big.Int).Div(raw, decimals)
		frac := new(big.Int).Mod(raw, decimals)
		return fmt.Sprintf("%s.%06d", whole.String(), frac.Int64())
	}

	log.Println("\n=== V2 Balances ===")
	log.Printf("pUSD Balance:   %s (%s raw)", fmtBal(info.PUSDBalance), info.PUSDBalance.String())
	log.Printf("USDC Balance:   %s (%s raw)", fmtBal(info.USDCBalance), info.USDCBalance.String())
	log.Printf("USDC.e Balance: %s (%s raw)", fmtBal(info.USDCEBalance), info.USDCEBalance.String())

	log.Println("\n=== USDC Allowances ===")
	log.Printf("CollateralOnramp:            %s", info.USDCAllowanceOnramp.String())

	log.Println("\n=== USDC.e Allowances ===")
	log.Printf("CollateralOnramp:            %s", info.USDCEAllowanceOnramp.String())

	log.Println("\n=== pUSD Allowances ===")
	log.Printf("CtfCollateralAdapter:        %s", info.PUSDAllowanceCtfAdapter.String())
	log.Printf("NegRiskCtfCollateralAdapter: %s", info.PUSDAllowanceNegRiskCtfAdapter.String())
	log.Printf("CollateralOfframp:           %s", info.PUSDAllowanceOfframp.String())

	log.Println("\n=== CTF Approvals ===")
	log.Printf("ExchangeV2:                  %v", info.CTFApprovedExchangeV2)
	log.Printf("NegRiskExchangeV2:           %v", info.CTFApprovedNegRiskExchangeV2)
	log.Printf("CtfCollateralAdapter:        %v", info.CTFApprovedCtfCollateralAdapter)
	log.Printf("NegRiskCtfCollateralAdapter: %v", info.CTFApprovedNegRiskCtfCollateralAdapter)
}
