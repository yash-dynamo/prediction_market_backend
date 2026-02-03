package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"prediction_market_backend/constants"
	"prediction_market_backend/go_bindings"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum network")

	usdcAddress := constants.MockUSDCAddress
	marketAddress := constants.PredictionMarketAddress

	usdc, err := go_bindings.NewMockUSDC(common.HexToAddress(usdcAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = marketAddress, usdc

	fmt.Println("contracts loaded successfully")
}
