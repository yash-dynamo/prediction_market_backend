package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"prediction_market_backend/go_bindings"
)

func main() {
	
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum network")

	usdcAddress := ""
    marketAddress := ""

	usdc, err := go_bindings.NewMockUSDC(common.HexToAddress(usdcAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contracts loaded successfully")
}