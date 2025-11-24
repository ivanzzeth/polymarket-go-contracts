package main

import (
	"context"
	"log"
	"math/big"
	"os"

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
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Parse private key from environment
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse privateKey: %v", err)
	}

	// Get address from private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("Deployer Address: %s", address.Hex())

	privateKeySigner, err := signer.NewSafeTradingPrivateKeySigner(chainID, client, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize V2 contract interface
	config := polymarketcontracts.GetContractConfig(polygonChainID)
	polymarketInterface, err := polymarketcontracts.NewContractInterface(client, polymarketcontracts.WithContractConfig(config), polymarketcontracts.WithSafeSigner(privateKeySigner))
	if err != nil {
		log.Fatalf("Failed to create contract interface V2: %v", err)
	}

	log.Printf("Safe Factory Address: %s", config.SafeProxyFactory.Hex())

	// Deploy Safe contract using V2
	log.Println("Deploying Safe contract...")
	safeProxy, txHash, err := polymarketInterface.DeploySafe()
	if err != nil {
		log.Fatalf("Failed to deploy Safe: %v", err)
	}

	log.Printf("✅ Safe deployed successfully!")
	log.Printf("   Transaction Hash: %s", txHash.Hex())
	log.Printf("   Safe Proxy Address: %s", safeProxy.Hex())
	log.Printf("   Owner Address: %s", address.Hex())

	// Verify the Safe address
	computedSafe, err := polymarketInterface.GetSafeAddress(address)
	if err != nil {
		log.Printf("Warning: Failed to compute Safe address: %v", err)
	} else {
		log.Printf("   Computed Safe Address: %s", computedSafe.Hex())
		if computedSafe.Hex() == safeProxy.Hex() {
			log.Println("   ✅ Address verification successful!")
		} else {
			log.Println("   ⚠️  Address mismatch - this may be expected depending on the factory implementation")
		}
	}

	log.Println("\nYou can now use this Safe address for multi-signature operations.")
	log.Printf("View Safe: https://evm-explorer.web3gate.xyz/evm/137/address/%s", safeProxy.Hex())
}
