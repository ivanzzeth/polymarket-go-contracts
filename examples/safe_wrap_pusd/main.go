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

	amountStr := os.Getenv("AMOUNT")
	if amountStr == "" {
		log.Fatal("AMOUNT environment variable is required in raw units (e.g., '1000000' for 1 token)")
	}
	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Fatalf("Invalid AMOUNT value: %s", amountStr)
	}

	asset := config.Collateral
	if assetHex := os.Getenv("ASSET"); assetHex != "" {
		asset = common.HexToAddress(assetHex)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	log.Println("\n=== Wrapping to pUSD via Safe ===")
	log.Printf("Asset (source):  %s", asset.Hex())
	log.Printf("Recipient:       %s", safeAddr.Hex())
	log.Printf("Amount:          %s tokens (%s human)", amount.String(), amountStr)

	txHash, err := v2.WrapToPUSDForSafe(ctx, safeSigner, polygonChainID, asset, safeAddr, amount)
	if err != nil {
		log.Fatalf("Failed to wrap to pUSD: %v", err)
	}

	log.Println("\nWrap successful!")
	log.Printf("Transaction: %s", txHash.Hex())
	log.Printf("https://polygonscan.com/tx/%s", txHash.Hex())
}
