package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

type activation struct {
	blockchain                     string
	address                        string
	passphrase                     string
	amountOfTransactionsToActivate uint64
	assetsToActivate               []enulibgo.RippleAsset
	comment                        string
}

var activations = []activation{
	{blockchain: "counterparty",
		address:                        "13sPWm9hPEJWb3fUCZ3TEHsnEB4ALagUfN",
		passphrase:                     "",
		amountOfTransactionsToActivate: 100,
		assetsToActivate:               nil,
		comment:                        "Passphrase should be empty for Counterparty because it isn't necessary to activate a wallet to receive certain assets"},
	{blockchain: "ripple",
		address:                        "rNx9Ddu8Hgi34GdSKhdHu9fuojH9Cb1gv2",
		passphrase:                     "",
		amountOfTransactionsToActivate: 100,
		assetsToActivate:               nil,
		comment:                        "A passphrase is not required for Ripple when no assets to enable is specified"},
	{blockchain: "ripple", address: "rst9JK6ANWuVHMqimwVukUXWzfc532vBg9",
		passphrase:                     "this is the twelve word pass phrase that is the wallet seed",
		amountOfTransactionsToActivate: 100,
		assetsToActivate: []enulibgo.RippleAsset{
			{Currency: "GOLD", Issuer: "rG8bvoqzkNow8ttQkLY52v97qe3YKnx8v2"},
			{Currency: "VIT", Issuer: "rM8X4RXK8GWcqYS1LbFscTnei2Py55WHRQ"},
			{Currency: "LEET", Issuer: "rM8X4RXK8GWcqYS1LbFscTnei2Py55WHRQ"},
		},
		comment: "Activate address for 100 transactions and allow the assets GOLD, VIT and LEET. Note issuer is mandatory for assets on Ripple"},
}

func main() {
	enulibgo.Init()

	for _, a := range activations {
		log.Printf("Activating address %s\n", a.address)
		returnValue, err := enulibgo.ActivateAddress(a.blockchain, a.address, a.passphrase, a.amountOfTransactionsToActivate, a.assetsToActivate)
		if err != nil {
			log.Printf("Failed to activate address. (Request Id: %s)", returnValue.RequestId)
			log.Println(err.Error())
		}
		log.Printf("%+#v\n", returnValue)
	}
}
