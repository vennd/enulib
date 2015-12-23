package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

type blockchainWallet struct {
	Blockchain       string
	AddresesToCreate uint64
}

var blockchains = []blockchainWallet{
	{"counterparty", 20},
	{"ripple", 0},
	{"gammagammazero", 14},
}

func main() {
	enulibgo.Init()

	for _, blockchain := range blockchains {
		log.Printf("Creating %s wallet and requesting %d addresses\n", blockchain.Blockchain, blockchain.AddresesToCreate)
		wallet, err := enulibgo.CreateWallet(blockchain.Blockchain, blockchain.AddresesToCreate)

		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("%s wallet creation successful. Wallet passphrase created: %s (Request Id: %s)", blockchain.Blockchain, wallet.Passphrase, wallet.RequestId)

			for i, address := range wallet.Addresses {
				log.Printf("Wallet address %d: %s\n", i, address)
			}
		}

		log.Println()
	}
}
