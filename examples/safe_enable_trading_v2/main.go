package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
	"github.com/ivanzzeth/polymarket-go-contracts/v2/signer"
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

	// Verify Safe is deployed (V2 has its own GetSafeAddress method)
	safeAddr, err := v2.GetSafeAddress(eoaAddress)
	if err != nil {
		log.Fatalf("Failed to compute Safe address: %v", err)
	}
	log.Printf("Safe Address: %s", safeAddr.Hex())

	code, err := client.CodeAt(context.Background(), safeAddr, nil)
	if err != nil {
		log.Fatalf("Failed to check Safe deployment: %v", err)
	}
	if len(code) == 0 {
		log.Fatalf("Safe not deployed at %s. Deploy it first.", safeAddr.Hex())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== Enabling V2 Trading for Safe ===")
	log.Println("This will:")
	log.Println("  - Approve USDC.e to CollateralOnramp (for wrap)")
	log.Println("  - Approve USDC (native) to CollateralOnramp (for wrap)")
	log.Println("  - Approve pUSD to CtfCollateralAdapter, NegRiskCtfCollateralAdapter, CollateralOfframp")
	log.Println("  - Set CTF approvals for ExchangeV2, NegRiskExchangeV2, and both adapters")

	txHashes, err := v2.EnableTradingForSafe(ctx, safeSigner, polygonChainID)
	if err != nil {
		log.Fatalf("Failed to enable V2 trading: %v", err)
	}

	log.Println("\nV2 trading enabled!")
	log.Printf("Transaction hashes (%d total):", len(txHashes))
	for i, h := range txHashes {
		log.Printf("  %d. %s", i+1, h.Hex())
		log.Printf("     https://polygonscan.com/tx/%s", h.Hex())
	}
}
