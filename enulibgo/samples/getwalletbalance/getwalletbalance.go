package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

var wallets = []struct {
	Address      string
	BlockchainId string
}{
	{"1GaZfh9VhxL4J8tBt2jrDvictZEKc8kcHx", "counterparty"},
	{"rUaCtgeUFnSSiqH9DbuibuMg4rMmpphvcY", "ripple"},
}

func main() {
	enulibgo.Init()

	// Retrieve from counterparty and ripple using the same API call

	for _, wallet := range wallets {
		log.Printf("Checking address %s on %s", wallet.Address, wallet.BlockchainId)
		balance, err := enulibgo.GetWalletBalance(wallet.BlockchainId, wallet.Address)
		if err != nil {
			log.Printf("Wallet balance check failed. Status code: %d (Request Id: %s)", err.ErrCode, balance.RequestId)
			log.Fatal(err)
		} else {
			log.Printf("Balance of address %s (Request Id: %s)\n", balance.Address, balance.RequestId)
			for _, amount := range balance.Balances {
				log.Printf("%d %s\n", amount.Quantity, amount.Asset)
			}

		}
		log.Println("")
	}
}
