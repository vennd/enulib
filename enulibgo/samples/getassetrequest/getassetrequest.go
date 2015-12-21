package main

import (
	"os"
	"log"
	"encoding/json"
	
	"github.com/vennd/enulib/enulibgo"
)

func usage() {
	println("Usage:\n")
	println("getassetrequest <assetRequestId>")
	println("")
	println("Where <assetRequestId> is the assetRequestId returned after performing an asset creation request.")
}

func main() {
	// One parameter must be specified which is the assetId
	if len(os.Args) != 2 {
		usage()
		
		os.Exit(1)
	}
	
	assetId := os.Args[1]

	enulibgo.Init()

	// Set values for asset
//	var assetId = "46e24f05d7b7bbffbe30f65a729b5823"

	asset, err := enulibgo.GetAsset(assetId)
	if err != nil {
		log.Printf("Failed to get asset. (Request Id: %s)", asset.RequestId)
		log.Fatal(err)
	}
	
	prettyPrint, errMarshal := json.MarshalIndent(asset, "", "  ")
	if errMarshal != nil {
		log.Panic(errMarshal)
	}
	
	log.Println(string(prettyPrint))
}
