package main

import (
	"github.com/vennd/enulib/enulibgo"
	"log"
)

func main() {
	enulibgo.Init()

	// Set values for asset
	var assetId = "cd91fb5593e5336453fb454041149a67"

	asset, err := enulibgo.GetAsset(assetId)

	if err != nil {
		log.Printf("Failed to get asset. (Request Id: %s)", asset.RequestId)
		log.Fatal(err)
	}

	log.Printf("Reply: %+#v", asset)
}
