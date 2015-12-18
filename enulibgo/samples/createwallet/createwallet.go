package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

var blockchains = []string{
	"counterparty",
	"ripple",
	"gammagammazero",
}

func main() {
	enulibgo.Init()

	for _, blockchain := range blockchains {
		log.Printf("Creating %s wallet", blockchain)
		wallet, err := enulibgo.CreateWallet(blockchain)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			log.Printf("%s wallet creation successful. Wallet passphrase created: %s (Request Id: %s)", blockchain, wallet.Passphrase, wallet.RequestId)

			for i, address := range wallet.Addresses {
				log.Printf("Wallet address %d: %s\n", i, address)

			}
		}

		log.Println()
	}
}
