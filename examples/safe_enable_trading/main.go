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
	polymarketInterface, err := polymarketcontracts.NewContractInterface(client, polymarketcontracts.WithContractConfig(config), polymarketcontracts.WithSafeSigner(safeSigner))
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

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== Enabling Trading for Safe ===")

	// Enable trading through Safe
	// The function will internally create TransactOpts and handle gas estimation
	txHashes, err := polymarketInterface.EnableTrading(ctx)
	if err != nil {
		log.Fatalf("Failed to enable trading for Safe: %v", err)
	}

	log.Println("✅ Trading successfully enabled for Safe!")
	log.Printf("Safe can now trade on Polymarket: %s", safeAddr.Hex())
	log.Printf("Transaction hashes:")
	for i, txHash := range txHashes {
		log.Printf("  %d. %s", i+1, txHash.Hex())
		log.Printf("     View transaction: https://evm-explorer.web3gate.xyz/evm/137/tx/%s", txHash.Hex())
	}
}
