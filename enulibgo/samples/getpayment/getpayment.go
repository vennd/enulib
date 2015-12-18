package main

import (
	"github.com/vennd/enulib/enulibgo"
	"log"
)

func main() {
	enulibgo.Init()

	// Retrieve a payment based upon its unique paymentId
	var paymentId = "61f79bed6ce762987b129bb659c8255c"
	log.Printf("Retrieving payment for paymentId: %s\n", paymentId)
	payment, err := enulibgo.GetPayment(paymentId)
	if err != nil {
		log.Printf("Failed to get payment. Request Id: %s)", payment.RequestId)
		log.Fatal(err.Error())
	}

	log.Printf("Payment: %+#v\n", payment)
	log.Println()

	// Retrieves all payments for a particular address
	var address = "rpu8gxvRzQ2JLQMN7Goxs6x9zffH3sjQBd"
	log.Printf("Retrieving payments for address: %s\n", address)
	payments, err := enulibgo.GetPaymentByAddress(address)
	if err != nil {
		log.Printf("Failed to get payment. Request Id: %s)", payment.RequestId)
		log.Fatal(err)
	}

	for _, payment := range payments {
		log.Printf("Payment: %+#v", payment)
	}
}
