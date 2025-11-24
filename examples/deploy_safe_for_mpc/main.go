package main

import (
	"log"
	"math/big"
	"os"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarket "github.com/ivanzzeth/polymarket-go-contracts"
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
	log.Printf("Deployer Address (MPC): %s", address.Hex())

	// Initialize V2 contract interface
	config := polymarket.GetContractConfig(polygonChainID)
	polymarketInterface, err := polymarket.NewContractInterface(client, config)
	if err != nil {
		log.Fatalf("Failed to create contract interface V2: %v", err)
	}

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

	// Get TransactionSender for MPC
	txSender, err := polymarketInterface.GetTransactionSenderrByCoboMpcTransactionSender(coboMpcSigner)
	if err != nil {
		log.Fatalf("Failed to create MPC transaction sender: %v", err)
	}

	log.Printf("Safe Factory Address: %s", config.SafeProxyFactory.Hex())

	// Deploy Safe contract using MPC
	log.Println("Deploying Safe contract via Cobo MPC...")
	safeProxy, txHash, err := polymarketInterface.DeploySafe(txSender, coboMpcSigner)
	if err != nil {
		log.Fatalf("Failed to deploy Safe: %v", err)
	}

	log.Printf("✅ Safe deployed successfully via MPC!")
	log.Printf("   Transaction Hash: %s", txHash.Hex())
	log.Printf("   Safe Proxy Address: %s", safeProxy.Hex())
	log.Printf("   Owner Address (MPC): %s", address.Hex())

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

	log.Println("\n✅ Safe deployment via Cobo MPC completed successfully!")
	log.Println("You can now use this Safe address for multi-signature operations with MPC.")
	log.Printf("View on Polygonscan: https://polygonscan.com/address/%s", safeProxy.Hex())

	log.Println("\n📝 Next steps:")
	log.Println("   1. Fund the Safe with USDC for trading")
	log.Println("   2. Use EnableTradingForSafe to set up trading allowances")
	log.Println("   3. Start trading on Polymarket!")
}
