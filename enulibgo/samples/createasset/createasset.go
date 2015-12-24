package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

var assets = []struct {
	BlockchainId                  string
	IssuerAddress                 string
	Passphrase                    string
	DistributionAddress           string
	DistributionAddressPassphrase string
	Asset                         string
	Quantity                      uint64
	Divisible                     bool
	Comment                       string
}{
	{"counterparty", "1GaZfh9VhxL4J8tBt2jrDvictZEKc8kcHx", "this is the twelve word pass phrase that is the wallet seed", "", "", "EXPERIMENT", 300000, true, "Create a divisible Counterparty asset"},
	{"ripple", "rM8X4RXK8GWcqYS1LbFscTnei2Py55WHRQ", "this is a different twelve word passphrase that is the wallet seed", "", "", "RUBIX", 100000000000000000, true, "Create a billion 1,000,000,000 RUBIX on Ripple. The API will generate a distribution address for us"},
	{"ripple", "rM8X4RXK8GWcqYS1LbFscTnei2Py55WHRQ", "this is a different twelve word passphrase that is the wallet seed", "rNx9Ddu8Hgi34GdSKhdHu9fuojH9Cb1gv2", "this yet another different twelve word passphrase that is the wallet seed", "FLDC2", 100000000000000, true, "Create 1,000,000 FLDC2 on Ripple and we specify a distribution address"},
}

func main() {
	enulibgo.Init()

	for _, asset := range assets {
		log.Printf("Creating asset on %s %d of %s from issuer %s. Divisible: %t. Distribution address: %s\n",
			asset.BlockchainId, asset.Quantity, asset.Asset, asset.IssuerAddress, asset.Divisible, asset.DistributionAddress)
		result, err := enulibgo.CreateAsset(asset.Passphrase, asset.IssuerAddress, asset.DistributionAddress, asset.DistributionAddressPassphrase, asset.Asset, asset.Quantity, asset.Divisible, asset.BlockchainId)
		if err != nil {
			log.Println("Asset creation failed")
			log.Fatal(err.Error())
		}

		log.Printf("Asset creation request successful. The asset blockchain name is: %s, AssetId: %s, RequestId: %s\n", result.Asset, result.AssetId, result.RequestId)
	}
}
