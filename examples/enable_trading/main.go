package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ivanzzeth/ethsig"
	polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
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
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("EOA Address: %s", address.Hex())

	// Create EOA signer
	eoaSigner := ethsig.NewEthPrivateKeySigner(privateKey)

	// Initialize contract interface with EOA signer
	config := polymarketcontracts.GetContractConfig(polygonChainID)
	polymarketInterface, err := polymarketcontracts.NewContractInterface(client, polymarketcontracts.WithContractConfig(config), polymarketcontracts.WithEOASigner(eoaSigner))
	if err != nil {
		log.Fatalf("Failed to create contract interface: %v", err)
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== Enabling Trading for EOA ===")

	// Enable trading for EOA
	txHashes, err := polymarketInterface.EnableTrading(ctx)
	if err != nil {
		log.Fatalf("Failed to enable trading: %v", err)
	}

	log.Println("\n✅ Trading successfully enabled for EOA!")
	log.Printf("EOA Address: %s", address.Hex())
	log.Printf("Transaction hashes:")
	for i, txHash := range txHashes {
		log.Printf("  %d. %s", i+1, txHash.Hex())
		log.Printf("     View on Polygonscan: https://polygonscan.com/tx/%s", txHash.Hex())
	}
}
