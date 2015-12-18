package main

import (
	"log"

	"github.com/vennd/enulib/enulibgo"
)

var payments = []struct {
	BlockchainId       string
	SourceAddress      string
	DestinationAddress string
	Asset              string
	Issuer             string
	Quantity           uint64
	Passphrase         string
	Comment            string
}{
	{"counterparty", "1HdnKzzCKFzNEJbmYoa3RcY4MhKPP3NB7p", "1HpkZBjNFRFagyj6Q2adRSagkfNDERZhg1", "SHIMA", "", 1001, "this is the twelve word pass phrase that is the wallet seed"}, "Sending a counterparty asset",
	{"ripple", "rNx9Ddu8Hgi34GdSKhdHu9fuojH9Cb1gv2", "rpu8gxvRzQ2JLQMN7Goxs6x9zffH3sjQBd", "XRP", "", 1000000, "this is a different twelve word passphrase that is the wallet seed"},
	{"ripple", "rG8bvoqzkNow8ttQkLY52v97qe3YKnx8v2", "rst9JK6ANWuVHMqimwVukUXWzfc532vBg9", "GOLD", "rG8bvoqzkNow8ttQkLY52v97qe3YKnx8v2", 30050000000, "this yet another different twelve word passphrase that is the wallet seed", "Sends 300.5 GOLD issued by rG8bvoqzkNow8ttQkLY52v97qe3YKnx8v2 from rG8bvoqzkNow8ttQkLY52v97qe3YKnx8v2 to rst9JK6ANWuVHMqimwVukUXWzfc532vBg9"},
}

func main() {
	enulibgo.Init()

	for _, payment := range payments {
		log.Printf("Creating payment on blockchain %s from %s to %s of %d %s.%s\n", payment.BlockchainId, payment.SourceAddress, payment.DestinationAddress, payment.Quantity, payment.Issuer, payment.Asset)
		result, err := enulibgo.CreatePayment(payment.BlockchainId, payment.SourceAddress, payment.DestinationAddress, payment.Asset, payment.Issuer, payment.Quantity, payment.Passphrase)
		if err != nil {
			log.Println("Payment creation failed")
			log.Fatal(err.Error())
		}

		log.Printf("Payment creation successful. PaymentId: %s, RequestId: %s\n", result.PaymentId, result.RequestId)
	}

}
