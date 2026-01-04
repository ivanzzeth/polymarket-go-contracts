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

	amountStr := os.Getenv("AMOUNT")
	if amountStr == "" {
		log.Fatal("AMOUNT environment variable is required (in USDC, e.g., '100' for 100 USDC)")
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Fatalf("Invalid AMOUNT value: %s", amountStr)
	}

	// Convert to USDC units (6 decimals)
	usdcDecimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	amount.Mul(amount, usdcDecimals)

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== Splitting Position via Safe ===")
	log.Printf("Condition ID: %s", conditionId.Hex())
	log.Printf("Amount: %s USDC", amountStr)
	log.Println("\nThis will split your USDC into conditional tokens (outcome tokens).")
	log.Println("Binary market partition [1, 2] for Yes/No outcomes is handled automatically.")

	// Split position through Safe
	ctf := polymarketInterface.GetConditionalTokens()
	collectionId, err := ctf.GetCollectionId(nil, [32]byte{}, conditionId, big.NewInt(0))
	if err != nil {
		log.Fatalf("Failed to query collectionID: %v", err)
	}

	posId, err := ctf.GetPositionId(nil, polymarketInterface.GetConfig().Collateral, collectionId)
	if err != nil {
		log.Fatalf("Failed to query positionID: %v", err)
	}

	balance, err := ctf.BalanceOf(nil, safeAddr, posId)
	if err != nil {
		log.Fatalf("Failed to query balance0: %v", err)
	}

	log.Printf("Balance: %v", balance.String())

	txHash, err := polymarketInterface.Split(ctx, conditionId, amount)
	if err != nil {
		log.Fatalf("Failed to split position: %v", err)
	}

	log.Println("\n✅ Position split successfully!")
	log.Printf("Safe Address: %s", safeAddr.Hex())
	log.Printf("Transaction Hash: %s", txHash.Hex())
	log.Printf("View transaction: https://evm-explorer.web3gate.xyz/evm/137/tx/%s", txHash.Hex())

	log.Println("\n📝 What happened:")
	log.Printf("   • %s USDC was locked in the CTF contract", amountStr)
	log.Println("   • You received conditional tokens for each outcome")
	log.Println("   • These tokens can be traded on Polymarket")
	log.Println("\n💡 Next steps:")
	log.Println("   1. Trade your conditional tokens on Polymarket")
	log.Println("   2. After market resolution, redeem winning tokens for USDC")
	log.Println("   3. Or merge tokens back to USDC if you have all outcomes")
}
