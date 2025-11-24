package main

import (
	"context"
	"log"
	"math/big"
	"os"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

func main() {
	// Use Polygon mainnet
	polygonChainID := big.NewInt(137)
	coboChainId := "MATIC" // Cobo chain ID for Polygon

	// Connect to Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to dial ethclient with rpc %v: %v", os.Getenv("RPC_URL"), err)
	}

	// Get Cobo MPC configuration from environment
	coboPrivateKey := os.Getenv("COBO_PRIVATE_KEY")
	if coboPrivateKey == "" {
		log.Fatal("COBO_PRIVATE_KEY environment variable is required")
	}

	coboWalletId := os.Getenv("COBO_WALLET_ID")
	if coboWalletId == "" {
		log.Fatal("COBO_WALLET_ID environment variable is required")
	}

	coboAddress := os.Getenv("COBO_ADDRESS")
	if coboAddress == "" {
		log.Fatal("COBO_ADDRESS environment variable is required (the MPC wallet address)")
	}

	// Parse Cobo environment (default to dev)
	coboEnv := coboWaas2.DevEnv
	if os.Getenv("COBO_ENV") == "prod" {
		coboEnv = coboWaas2.ProdEnv
	}

	address := common.HexToAddress(coboAddress)
	log.Printf("MPC EOA Address: %s", address.Hex())

	// Create Cobo MPC signer
	coboMpcSigner, err := signer.NewCoboMpcSigner(
		coboEnv,
		coboPrivateKey,
		coboChainId,
		polygonChainID,
		coboWalletId,
		address,
	)
	if err != nil {
		log.Fatalf("Failed to create Cobo MPC signer: %v", err)
	}

	mpcSafeSigner, err := signer.NewSafeTradingSingleMpcSigner(polygonChainID, client, coboMpcSigner)
	if err != nil {
		log.Fatalf("Failed to create Cobo MPC safe signer: %v", err)
	}

	// Initialize contract interface
	config := polymarketcontracts.GetContractConfig(polygonChainID)
	polymarketInterface, err := polymarketcontracts.NewContractInterface(client, polymarketcontracts.WithContractConfig(config), polymarketcontracts.WithSafeSigner(mpcSafeSigner))
	if err != nil {
		log.Fatalf("Failed to create contract interface: %v", err)
	}

	// Calculate Safe address
	safeAddr, err := polymarketInterface.GetSafeAddress(address)
	if err != nil {
		log.Fatalf("Failed to get Safe address: %v", err)
	}

	log.Printf("Safe Address: %s", safeAddr.Hex())
	log.Printf("Safe Factory: %s", config.SafeProxyFactory.Hex())

	// Check if Safe is deployed
	code, err := client.CodeAt(context.Background(), safeAddr, nil)
	if err != nil {
		log.Fatalf("Failed to check Safe deployment: %v", err)
	}

	if len(code) == 0 {
		log.Printf("⚠️  Safe not deployed yet at %s", safeAddr.Hex())
		log.Println("Please deploy the Safe first using deploy_safe_for_mpc example")
		return
	}

	log.Println("✅ Safe is deployed")
	log.Println()

	// Create context with timeout
	ctx := context.Background()

	// Enable trading for Safe
	log.Println("=== Enabling Trading for MPC Safe ===")
	log.Println("This will set up all required allowances for trading on Polymarket")
	log.Println()

	txHashes, err := polymarketInterface.EnableTrading(ctx)
	if err != nil {
		log.Fatalf("Failed to enable trading: %v", err)
	}

	log.Println()
	log.Println("🎉 Success! Trading is now enabled for your MPC Safe wallet!")
	log.Println()
	log.Println("📝 What was configured:")
	log.Println("   ✅ USDC approved for Exchange contract")
	log.Println("   ✅ USDC approved for NegRiskAdapter contract")
	log.Println("   ✅ USDC approved for NegRiskExchange contract")
	log.Println("   ✅ CTF approved for Exchange contract")
	log.Println("   ✅ CTF approved for NegRiskAdapter contract")
	log.Println("   ✅ CTF approved for NegRiskExchange contract")
	log.Println()
	log.Printf("📝 Transaction hashes:")
	for i, txHash := range txHashes {
		log.Printf("   %d. %s", i+1, txHash.Hex())
		log.Printf("      View transaction: https://evm-explorer.web3gate.xyz/evm/137/tx/%s", txHash.Hex())
	}
	log.Println()
	log.Println("🚀 Next steps:")
	log.Println("   1. Fund your Safe with USDC for trading")
	log.Println("   2. Start trading on Polymarket!")
	log.Printf("   3. View your Safe: https://evm-explorer.web3gate.xyz/evm/137/address/%s", safeAddr.Hex())
}
