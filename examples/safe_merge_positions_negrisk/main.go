package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

func main() {
	// Use Polygon mainnet
	polygonChainID := big.NewInt(137)

	// Connect to Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to dial ethclient with rpc %v: %v", os.Getenv("RPC_URL"), err)
	}

	// Parse private key from environment
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse privateKey: %v", err)
	}

	// Get EOA address from private key
	eoaAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("EOA Address: %s", eoaAddress.Hex())

	// Create Safe trading signer from private key
	safeSigner, err := signer.NewSafeTradingPrivateKeySigner(polygonChainID, client, privateKey)
	if err != nil {
		log.Fatalf("Failed to create safe signer: %v", err)
	}

	// Initialize contract interface
	config := polymarketcontracts.GetContractConfig(polygonChainID)
	polymarketInterface, err := polymarketcontracts.NewContractInterface(
		client,
		polymarketcontracts.WithContractConfig(config),
		polymarketcontracts.WithSafeSigner(safeSigner),
	)
	if err != nil {
		log.Fatalf("Failed to create contract interface: %v", err)
	}

	// Compute Safe address for this EOA
	safeAddr, err := polymarketInterface.GetSafeAddress(eoaAddress)
	if err != nil {
		log.Fatalf("Failed to compute Safe address: %v", err)
	}
	log.Printf("Safe Address: %s", safeAddr.Hex())

	// Check if Safe is deployed
	code, err := client.CodeAt(context.Background(), safeAddr, nil)
	if err != nil {
		log.Fatalf("Failed to check Safe deployment: %v", err)
	}
	if len(code) == 0 {
		log.Fatalf("Safe not deployed at address %s. Please deploy the Safe first using the deploy_safe example.", safeAddr.Hex())
	}
	log.Println("✅ Safe is deployed")

	// Get parameters from environment
	conditionIdHex := os.Getenv("CONDITION_ID")
	if conditionIdHex == "" {
		log.Fatal("CONDITION_ID environment variable is required")
	}
	conditionId := common.HexToHash(conditionIdHex)
	amount := new(big.Int)

	amountStr := os.Getenv("AMOUNT")
	if amountStr != "" {
		var ok bool
		amount, ok = amount.SetString(amountStr, 10)
		if !ok {
			log.Fatalf("Invalid AMOUNT value: %s", amountStr)
		}
		// Convert to token units (6 decimals for CTF tokens)
		tokenDecimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
		amount.Mul(amount, tokenDecimals)
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Enable trading first (this will set up all necessary approvals)
	log.Println("\n=== Enabling Trading for Safe ===")
	txHashes, err := polymarketInterface.EnableTrading(ctx)
	if err != nil {
		log.Fatalf("Failed to enable trading for Safe: %v", err)
	}
	log.Println("✅ Trading enabled for Safe")
	if len(txHashes) > 0 {
		log.Printf("Approval transaction hashes:")
		for i, txHash := range txHashes {
			log.Printf("  %d. %s", i+1, txHash.Hex())
		}
	}

	log.Println("\n=== Merging NegRisk Positions via Safe ===")
	log.Printf("Condition ID: %s", conditionId.Hex())
	log.Printf("Amount: %s tokens", amountStr)
	log.Printf("EOA: %s", eoaAddress)
	log.Printf("Safe: %s", safeAddr)
	log.Println("\nThis will merge your conditional tokens back into USDC for a NegRisk market.")
	log.Println("⚠️  Note: This is for NegRisk markets only. Use safe_merge_positions for standard markets.")
	log.Println("⚠️  Note: You must hold ALL outcome tokens for this market to merge successfully!")

	// Merge positions through Safe using NegRisk adapter
	txHash, err := polymarketInterface.MergeNegRisk(ctx, conditionId, amount)
	if err != nil {
		log.Fatalf("Failed to merge NegRisk positions: %v", err)
	}

	log.Println("\n✅ NegRisk positions merged successfully!")
	log.Printf("Safe Address: %s", safeAddr.Hex())
	log.Printf("Transaction Hash: %s", txHash.Hex())
	log.Printf("View transaction: https://evm-explorer.web3gate.xyz/evm/137/tx/%s", txHash.Hex())

	log.Println("\n📝 What happened:")
	log.Printf("   • %s conditional tokens (all outcomes) were burned via NegRiskAdapter", amountStr)
	log.Printf("   • %s USDC was unlocked from the CTF contract", amountStr)
	log.Println("   • USDC is now available in your Safe wallet")
	log.Println("\n💡 Use cases for merge:")
	log.Println("   • Cancel a position before trading")
	log.Println("   • Collect full sets from trading and convert back to USDC")
	log.Println("   • Exit a market position by acquiring all outcome tokens")
}
