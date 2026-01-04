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

	ctf := polymarketInterface.GetConditionalTokens()
	collectionId, err := ctf.GetCollectionId(nil, [32]byte{}, conditionId, big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to query collectionID: %v", err)
	}

	collectionId1, err := ctf.GetCollectionId(nil, [32]byte{}, conditionId, big.NewInt(2))
	if err != nil {
		log.Fatalf("Failed to query collectionID: %v", err)
	}

	posId, err := ctf.GetPositionId(nil, polymarketInterface.GetConfig().Collateral, collectionId)
	if err != nil {
		log.Fatalf("Failed to query positionID: %v", err)
	}
	posId1, err := ctf.GetPositionId(nil, polymarketInterface.GetConfig().Collateral, collectionId1)
	if err != nil {
		log.Fatalf("Failed to query positionID: %v", err)
	}

	balance, err := ctf.BalanceOf(nil, safeAddr, posId)
	if err != nil {
		log.Fatalf("Failed to query balance0: %v", err)
	}
	balance1, err := ctf.BalanceOf(nil, safeAddr, posId1)
	if err != nil {
		log.Fatalf("Failed to query balance0: %v", err)
	}

	if amount.Cmp(common.Big0) == 0 {
		amount = balance
	}

	log.Println("\n=== Merging Positions via Safe ===")
	log.Printf("Condition ID: %s", conditionId.Hex())
	log.Printf("Amount: %s tokens", amountStr)
	log.Printf("EOA: %s", eoaAddress)
	log.Printf("Funder: %s", safeAddr)
	log.Printf("posId: %s", posId)
	log.Printf("posId1: %s", posId1)

	log.Printf("Position Balance: %v", balance.String())
	log.Printf("Position Balance1: %v", balance1.String())

	log.Println("\nThis will merge your conditional tokens back into USDC.")
	log.Println("⚠️  Note: You must hold ALL outcome tokens for this market to merge successfully!")

	// Merge positions through Safe
	txHash, err := polymarketInterface.Merge(ctx, conditionId, amount)
	if err != nil {
		log.Fatalf("Failed to merge positions: %v", err)
	}

	log.Println("\n✅ Positions merged successfully!")
	log.Printf("Safe Address: %s", safeAddr.Hex())
	log.Printf("Transaction Hash: %s", txHash.Hex())
	log.Printf("View transaction: https://evm-explorer.web3gate.xyz/evm/137/tx/%s", txHash.Hex())

	log.Println("\n📝 What happened:")
	log.Printf("   • %s conditional tokens (all outcomes) were burned", amountStr)
	log.Printf("   • %s USDC was unlocked from the CTF contract", amountStr)
	log.Println("   • USDC is now available in your Safe wallet")
	log.Println("\n💡 Use cases for merge:")
	log.Println("   • Cancel a position before trading")
	log.Println("   • Collect full sets from trading and convert back to USDC")
	log.Println("   • Exit a market position by acquiring all outcome tokens")
}
