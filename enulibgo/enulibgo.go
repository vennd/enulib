package enulibgo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var enuAPIBaseURL = "https://enu.io"

//var enuBetaAPIBaseURL = "https://beta.enu.io"
//var enuAPITestBaseURL = "http://localhost:8080"

var ApiKey string
var ApiSecret string

func Init() {
	var license License

	// Read license key from file
	file, err := ioutil.ReadFile("./enu_key.json")
	if err != nil {
		log.Println("Unable to read API key")
		log.Fatalln(err)
	}

	err = json.Unmarshal(file, &license)

	ApiKey = license.Key
	ApiSecret = license.Secret
}

func baseURL() string {
	urlOverride := os.Getenv("ENU_BASE_URL")

	if urlOverride == "" {
		return enuAPIBaseURL
	}

	_, err := url.Parse(urlOverride)
	if err != nil {
		return enuAPIBaseURL
	}

	return urlOverride
}

func ComputeHmac512(message []byte, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

func GetNonce() int64 {
	return time.Now().Unix()
}

func DoEnuAPI(method string, url string, postData []byte) ([]byte, *EnuError) {
	if method != "POST" && method != "GET" {
		return nil, &EnuError{ErrCode: 13, Err: errors.New("DoEnuAPI can only be called for a POST or GET method")}
	}

	postDataJson := string(postData)

	//	log.Printf("Posting to: %s", url)
	//	log.Printf("Posting: %s", postDataJson)

	// Set headers
	req, err := http.NewRequest(method, url, bytes.NewBufferString(postDataJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accessKey", ApiKey)
	req.Header.Set("signature", ComputeHmac512(postData, ApiSecret))

	// Perform request
	clientPointer := &http.Client{}
	resp, err := clientPointer.Do(req)
	if err != nil {
		return nil, &EnuError{ErrCode: 13, Err: err}
	}

	// Did not receive an OK or Accepted
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		message := fmt.Sprintf("Request failed. Status code: %d\n", resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, &EnuError{ErrCode: 13, Err: errors.New(message)}
		}

		// Expect to be able to unmarshall a returnerror structure
		var returnError ReturnError
		if err := json.Unmarshal(body, &returnError); err != nil {
			return body, &EnuError{ErrCode: 13, Err: errors.New(string(body)), RequestId: "n/a"}
		}

		return body, &EnuError{ErrCode: returnError.Code, Err: errors.New(returnError.Description), RequestId: returnError.RequestId}
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return body, &EnuError{ErrCode: 13, Err: err}
	}

	log.Printf("Reply: %s\n", string(body))

	return body, nil
}

func ActivateAddress(blockchainId string, address string, passphrase string, amountOfTransactionsToActivate uint64, assetsToActivate []RippleAsset) (ActivateAddressStruct, *EnuError) {
	var activation ActivateAddressStruct

	// Make URL from base URL
	var url = baseURL() + "/wallet/activate/address/" + address

	type genericMap map[string]interface{}

	var send = map[string]interface{}{
		"address": address,
		"amount":  amountOfTransactionsToActivate,
	}

	if passphrase != "" {
		send["passphrase"] = passphrase
	}

	if assetsToActivate != nil {
		send["assets"] = assetsToActivate
	}

	jsonBytes, err := json.Marshal(send)
	if err != nil {
		return activation, &EnuError{ErrCode: 13, Err: err}
	}

	responseData, err := DoEnuAPI("POST", url, jsonBytes)
	if err.(*EnuError) != nil {
		return activation, &EnuError{ErrCode: 13, Err: err}
	}

	if err := json.Unmarshal(responseData, &activation); err != nil {
		return activation, &EnuError{ErrCode: 13, Err: err}
	}

	return activation, nil
}

// Creates a wallet on the specified blockchain
// blockchainId may be set to an empty string "" if you wish to use your default blockchain
func CreateWallet(blockchainId string) (Wallet, *EnuError) {
	var basicRequest BasicRequest
	var wallet Wallet

	// Set parameters
	var url = baseURL() + "/wallet"
	basicRequest.BlockchainId = blockchainId
	basicRequest.Nonce = GetNonce()

	// Marshall to json
	jsonBytes, err := json.Marshal(basicRequest)
	if err != nil {
		return wallet, &EnuError{ErrCode: 13, Err: err}
	}

	// Post to Enu server
	responseData, err := DoEnuAPI("POST", url, jsonBytes)
	if err.(*EnuError) != nil {
		return wallet, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	// Unmarshall from json
	if err := json.Unmarshal(responseData, &wallet); err != nil {
		return wallet, &EnuError{ErrCode: 13, Err: err}
	}

	return wallet, nil
}

func CreateAsset(passphrase string, sourceAddress string, distributionAddress string, distributionAddressPassphrase string, asset string, quantity uint64, divisible bool, blockchain string) (Asset, *EnuError) {
	var assetStruct Asset

	// Make URL from base URL
	var url = baseURL() + "/asset"
	var send = map[string]interface{}{
		"passphrase":             passphrase,
		"sourceAddress":          sourceAddress,
		"distributionAddress":    distributionAddress,
		"distributionPassphrase": distributionAddressPassphrase,
		"asset":                  asset,
		"quantity":               quantity,
		"divisible":              divisible,
		"nonce":                  GetNonce(),
		"blockchainId":           blockchain,
	}

	// Marshall to json
	assetJsonBytes, err := json.Marshal(send)
	if err != nil {
		return assetStruct, &EnuError{ErrCode: 13, Err: err}
	}

	// Post to Enu server
	responseData, err := DoEnuAPI("POST", url, assetJsonBytes)
	if err.(*EnuError) != nil {
		return assetStruct, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	if err := json.Unmarshal(responseData, &assetStruct); err != nil {
		return assetStruct, &EnuError{ErrCode: 13, Err: err}
	}

	return assetStruct, nil
}

func CreateDividend(passphrase string, sourceAddress string, asset string, dividendAsset string, quantityPerUnit uint64) (Dividend, *EnuError) {
	var dividendStruct Dividend
	// Make URL from base URL
	var url = baseURL() + "/asset/dividend"
	var send = map[string]interface{}{
		"passphrase":      passphrase,
		"sourceAddress":   sourceAddress,
		"asset":           asset,
		"dividendAsset":   dividendAsset,
		"quantityPerUnit": quantityPerUnit,
		"nonce":           GetNonce(),
	}

	dividendJsonBytes, err := json.Marshal(send)
	if err != nil {
		return dividendStruct, &EnuError{ErrCode: 13, Err: err}
	}

	responseData, err := DoEnuAPI("POST", url, dividendJsonBytes)
	if err.(*EnuError) != nil {
		return dividendStruct, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	if err := json.Unmarshal(responseData, &dividendStruct); err != nil {
		return dividendStruct, &EnuError{ErrCode: 13, Err: err}
	}

	return dividendStruct, nil

}

// Creates a payment from the sourceAddress to the destinationAddress
// Where the blockchain is "ripple" and a custom asset is being paid, an issuer must be provided
func CreatePayment(blockchain string, sourceAddress string, destinationAddress string, asset string, issuer string, quantity uint64, passphrase string) (Payment, *EnuError) {
	//	var paymentId string
	var paymentStruct Payment
	// Make URL from base URL
	var url = baseURL() + "/wallet/payment"

	type genericMap map[string]interface{}

	var send = map[string]interface{}{
		"blockchainId":       blockchain,
		"passphrase":         passphrase,
		"sourceAddress":      sourceAddress,
		"destinationAddress": destinationAddress,
		"asset":              asset,
		"issuer":             issuer,
		"quantity":           quantity,
		"nonce":              GetNonce(),
	}

	payloadJsonBytes, err := json.Marshal(send)
	if err != nil {
		return paymentStruct, &EnuError{ErrCode: 13, Err: err}
	}

	responseData, err := DoEnuAPI("POST", url, payloadJsonBytes)
	if err.(*EnuError) != nil {
		return paymentStruct, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	if err := json.Unmarshal(responseData, &paymentStruct); err != nil {
		return paymentStruct, &EnuError{ErrCode: 13, Err: err}
	}

	return paymentStruct, nil
}

// Retrieves the balance of all the assets held in a particular address
// blockchainId may be set to an empty string "" if you wish to use your default blockchain
func GetWalletBalance(blockchainId string, address string) (AddressBalances, *EnuError) {
	var result AddressBalances
	var basicRequest BasicRequest

	// Set parameters
	var url = baseURL() + "/wallet/balances/" + address
	basicRequest.BlockchainId = blockchainId
	basicRequest.Nonce = GetNonce()

	// Marshall to json
	jsonBytes, err := json.Marshal(basicRequest)
	if err != nil {
		return result, &EnuError{ErrCode: 13, Err: err}
	}

	// Send the request
	responseData, err := DoEnuAPI("GET", url, jsonBytes)
	if err.(*EnuError) != nil {
		return result, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	// Unmarshall the data
	if err := json.Unmarshal(responseData, &result); err != nil {
		return result, &EnuError{ErrCode: 13, Err: err}
	}

	return result, nil
}

func GetAssetLedger(asset string) (AssetBalances, *EnuError) {
	var balance AssetBalances

	// Make URL from base URL
	var url = baseURL() + "/asset/ledger/" + asset

	//var empty []byte
	balance.Nonce = GetNonce()
	payloadJsonBytes, err := json.Marshal(balance)

	if err != nil {
		return balance, &EnuError{ErrCode: 13, Err: err}
	}

	responseData, err := DoEnuAPI("GET", url, payloadJsonBytes)
	if err.(*EnuError) != nil {
		return balance, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	if err := json.Unmarshal(responseData, &balance); err != nil {
		return balance, &EnuError{ErrCode: 13, Err: err}
	}

	return balance, nil
}

func GetAssetIssuances(asset string) (AssetIssuances, *EnuError) {
	var assetIssuances AssetIssuances

	// Make URL from base URL
	var url = baseURL() + "/asset/issuances/" + asset

	//var empty []byte
	assetIssuances.Nonce = GetNonce()
	payloadJsonBytes, err := json.Marshal(assetIssuances)

	if err != nil {
		return assetIssuances, &EnuError{ErrCode: 13, Err: err}
	}

	responseData, err := DoEnuAPI("GET", url, payloadJsonBytes)
	if err.(*EnuError) != nil {
		return assetIssuances, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	log.Println(string(responseData))

	if err := json.Unmarshal(responseData, &assetIssuances); err != nil {
		return assetIssuances, &EnuError{ErrCode: 13, Err: err}
	}

	return assetIssuances, nil
}

// Retrieves a payment by it's unique paymentId
func GetPayment(paymentId string) (Payment, *EnuError) {
	var payment Payment

	// Make URL from base URL
	var url = baseURL() + "/wallet/payment/" + paymentId

	log.Println(url)

	// Make the JSON request
	var empty []byte = []byte("{}")
	payment.Nonce = GetNonce()

	// Send the request
	responseData, err := DoEnuAPI("GET", url, empty)
	if err != nil {
		return payment, &EnuError{ErrCode: err.ErrCode, Err: err.Err}
	}

	// Unmarshall the data
	if err := json.Unmarshal(responseData, &payment); err != nil {
		return payment, &EnuError{ErrCode: 13, Err: err}
	}

	return payment, nil
}

// Retrieves all payments that were received or sent from a particular address
func GetPaymentByAddress(address string) ([]Payment, *EnuError) {
	var payments []Payment

	// Make URL from base URL
	var url = baseURL() + "/payment/address/" + address

	log.Println(url)

	// Make the JSON request
	var empty []byte = []byte("{}")

	// Send the request
	responseData, err := DoEnuAPI("GET", url, empty)
	if err != nil {
		return payments, &EnuError{ErrCode: err.ErrCode, Err: err.Err}
	}

	// Unmarshall the data
	if err := json.Unmarshal(responseData, &payments); err != nil {
		return payments, &EnuError{ErrCode: 13, Err: err}
	}

	return payments, nil
}

func GetDividend(dividendId string) (Dividend, *EnuError) {
	var dividend Dividend

	// Make URL from base URL
	var url = baseURL() + "/asset/dividend/" + dividendId

	// Make the JSON request
	//var empty []byte
	dividend.Nonce = GetNonce()

	payloadJsonBytes, err := json.Marshal(dividend)
	if err != nil {
		return dividend, &EnuError{ErrCode: 13, Err: err}
	}

	// Send the request
	responseData, err := DoEnuAPI("GET", url, payloadJsonBytes)
	if err.(*EnuError) != nil {
		return dividend, &EnuError{ErrCode: err.(*EnuError).ErrCode, Err: err.(*EnuError).Err}
	}

	// Unmarshall the data
	if err := json.Unmarshal(responseData, &dividend); err != nil {
		return dividend, &EnuError{ErrCode: 13, Err: err}
	}

	return dividend, nil
}

func GetAsset(assetId string) (Asset, *EnuError) {
	var asset Asset

	// Make URL from base URL
	var url = baseURL() + "/asset/" + assetId

	// Make the JSON request
	//var empty []byte
	asset.Nonce = GetNonce()
	payloadJsonBytes, err := json.Marshal(asset)
	if err != nil {
		return asset, &EnuError{ErrCode: 13, Err: err}
	}

	// Send the request
	responseData, err := DoEnuAPI("GET", url, payloadJsonBytes)
	if err.(*EnuError) != nil {
		return asset, &EnuError{ErrCode: 13, Err: err}
	}

	// Unmarshall the data
	if err := json.Unmarshal(responseData, &asset); err != nil {
		return asset, &EnuError{ErrCode: 13, Err: err}
	}

	return asset, nil
}
