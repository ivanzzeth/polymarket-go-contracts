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
	v1, err := polymarketcontracts.NewContractInterface(
		client,
		polymarketcontracts.WithContractConfig(config),
		polymarketcontracts.WithSafeSigner(safeSigner),
	)
	if err != nil {
		log.Fatalf("Failed to create V1 contract interface: %v", err)
	}

	safeAddr, err := v1.GetSafeAddress(eoaAddress)
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

	v2, err := polymarketcontracts.NewContractInterfaceV2(
		client,
		config,
		safeSigner,
		v1.GetSafeAddress,
		v1.ExecuteTransactionBySafeAndSingleSigner,
	)
	if err != nil {
		log.Fatalf("Failed to create V2 contract interface: %v", err)
	}

	conditionIdHex := os.Getenv("CONDITION_ID")
	if conditionIdHex == "" {
		log.Fatal("CONDITION_ID environment variable is required")
	}
	conditionId := common.HexToHash(conditionIdHex)

	amountStr := os.Getenv("AMOUNT")
	if amountStr == "" {
		log.Fatal("AMOUNT environment variable is required in raw units (e.g., '1000000' for 1 token)")
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Fatalf("Invalid AMOUNT value: %s", amountStr)
	}

	partition := []*big.Int{big.NewInt(1), big.NewInt(2)}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== V2 Merge Positions via Safe (CtfCollateralAdapter) ===")
	log.Printf("Condition ID: %s", conditionId.Hex())
	log.Printf("Amount: %s tokens", amountStr)
	log.Println("Partition: [1, 2] (binary Yes/No)")
	log.Println("\nThis merges conditional tokens back into pUSD via the CtfCollateralAdapter.")
	log.Println("You must hold ALL outcome tokens to merge successfully.")

	var condID [32]byte
	copy(condID[:], conditionId[:])

	txHash, err := v2.MergePositionsForSafe(ctx, safeSigner, polygonChainID, condID, partition, amount)
	if err != nil {
		log.Fatalf("Failed to merge positions: %v", err)
	}

	log.Println("\nMerge successful!")
	log.Printf("Transaction: %s", txHash.Hex())
	log.Printf("https://polygonscan.com/tx/%s", txHash.Hex())
}
