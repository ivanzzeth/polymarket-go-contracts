package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	polymarket "github.com/ivanzzeth/polymarket-go-contracts"
)

func main() {
	polygonChainID := big.NewInt(137)
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to dial ethclien with rpc %v: %v", os.Getenv("RPC_URL"), err)
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse privateKey: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("Address: %v", address.Hex())

	polymarketInterface, err := polymarket.NewContractInterface(client, polymarket.GetContractConfig(polygonChainID))
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, polygonChainID)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	err = polymarketInterface.EnableTrading(ctx, auth, address)
	if err != nil {
		log.Fatalf("Failed to enable trading: %v", err)
	}
}
