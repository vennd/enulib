package enulibgo

import (
	"strconv"
)

type EnuError struct {
	ErrCode   int64
	Err       error
	RequestId string
}

type ReturnError struct {
	RequestId   string `json:"requestId"`
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

func (e *EnuError) Error() string {
	var result string

	if e == nil {
		return "<nil>"
	}

	if e.ErrCode != 0 {
		result += "Error Code: " + strconv.FormatInt(e.ErrCode, 10)
	}

	if e.Err != nil {
		result += ", Error Message: " + e.Err.Error()
	}

	return result
}

type License struct {
	Comment string `json:comment`
	Key     string `json:"key"`
	Secret  string `json:"secret"`
}

type Block struct {
	BlockId  int64  `json:"blockId"`
	Status   string `json:"status"`
	Duration int64  `json:"duration"`
}

type Blocks struct {
	Allblocks []Block `json:"blocks"`
	RequestId string  `json:"requestId`
	Nonce     int64   `json:"nonce"`
}

// General struct for requests which don't have any parameters
type BasicRequest struct {
	BlockchainId string `json:"blockchainId,omitempty"`
	Nonce        int64  `json:"nonce,omitempty"`
}

type Payment struct {
	BlockchainId            string `json:"blockchainId"`
	SourceAddress           string `json:"sourceAddress"`
	DestinationAddress      string `json:"destinationAddress"`
	Asset                   string `json:"asset"`
	Issuer                  string `json:"issuer"`
	Amount                  uint64 `json:"amount"`
	PaymentId               string `json:"paymentId"`
	TxFee                   int64  `json:"txFee"`
	BroadcastTxId           string `json:"broadcastTxId"`
	BlockchainStatus        string `json:"blockchainStatus"`
	BlockchainConfirmations uint64 `json:"blockchainConfirmations"`
	PaymentTag              string `json:"paymentTag"`
	Status                  string `json:"status"`
	ErrorCode               int64  `json:"errorCode"`
	ErrorMessage            string `json:"errorMessage"`
	RequestId               string `json:"requestId"`
	Nonce                   int64  `json:"nonce"`
}

type Wallet struct {
	Passphrase    string   `json:"passphrase"`
	HexSeed       string   `json:"hexSeed"`
	Addresses     []string `json:"addresses"`
	RequestId     string   `json:"requestId"`
	BlockchainId  string   `json:"blockchainId,omitempty"`
	KeyType       string   `json:"key_type,omitempty"`
	MasterKey     string   `json:"master_key,omitempty"`
	MasterSeed    string   `json:"master_seed,omitempty"`
	MasterSeedHex string   `json:"master_seed_hex,omitempty"`
	PublicKey     string   `json:"public_key,omitempty"`
	PublicKeyHex  string   `json:"public_key_hex,omitempty"`
}

type Address struct {
	Value      string `json:"value"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	RequestId  string `json:"requestId"`
	Nonce      int64  `json:"nonce"`
}

type Send struct {
	Passphrase         string `json:"passphrase"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	Asset              string `json:"asset"`
	Quantity           uint64 `json:"quantity"`
}

type Asset struct {
	Passphrase              string `json:"passphrase,omitempty"`
	SourceAddress           string `json:"sourceAddress"`
	DistributionPassphrase  string `json:"distributionPassphrase,omitempty"`
	DistributionAddress     string `json:"distributionAddress,omitempty"`
	AssetId                 string `json:"assetId"`
	Asset                   string `json:"asset"`
	Issuer                  string `json:"issuer,omitempty"`
	Description             string `json:"description"`
	Quantity                uint64 `json:"quantity"`
	Divisible               bool   `json:"divisible"`
	BroadcastTxId           string `json:"broadcastTxId"`
	BlockchainStatus        string `json:"blockchainStatus"`
	BlockchainConfirmations uint64 `json:"blockchainConfirmations"`
	Status                  string `json:"status"`
	ErrorMessage            string `json:"errorMessage"`
	RequestId               string `json:"requestId"`
	Nonce                   int64  `json:"nonce"`
	BlockchainId            string `json:"blockchainId"`
}

type Amount struct {
	Asset             string  `json:"asset"`
	Quantity          uint64  `json:"quantity"`
	PercentageHolding float64 `json:"percentageHolding"`
}

type AddressAmount struct {
	Address           string  `json:"address"`
	Quantity          uint64  `json:"quantity"`
	PercentageHolding float64 `json:"percentageHolding"`
}

type AddressBalances struct {
	Address   string   `json:"address"`
	Balances  []Amount `json:"balances"`
	RequestId string   `json:"requestId`
	Nonce     int64    `json:"nonce"`
}

type AssetBalances struct {
	Asset        string          `json:"asset"`
	Locked       bool            `json:"locked"`
	Divisible    bool            `json:"divisible"`
	Divisibility uint64          `json:"divisibility"`
	Description  string          `json:"description"`
	Supply       uint64          `json:"quantity"`
	Balances     []AddressAmount `json:"balances"`
	RequestId    string          `json:"requestId`
	Nonce        int64           `json:"nonce"`
}

type Issuance struct {
	BlockIndex uint64 `json:"block_index"`
	Quantity   uint64 `json:"quantity"`
	Issuer     string `json:"issuer"`
	Transfer   bool   `json:"transfer"`
}
type AssetIssuances struct {
	Asset        string     `json:"asset"`
	Divisible    bool       `json:"divisible"`
	Divisibility uint64     `json:"divisibility"`
	Description  string     `json:"description"`
	Locked       bool       `json:"locked"`
	Issuances    []Issuance `json:"issuances"`
	RequestId    string     `json:"requestId"`
	Nonce        int64      `json:"nonce"`
}

type Dividend struct {
	Passphrase      string `json:"passphrase"`
	SourceAddress   string `json:"sourceAddress"`
	DividendId      string `json:"dividendId"`
	Asset           string `json:"asset"`
	DividendAsset   string `json:"dividendAsset"`
	QuantityPerUnit uint64 `json:"quantityPerUnit"`
	Status          string `json:"status"`
	ErrorMessage    string `json:"errorMessage`
	RequestId       string `json:"requestId"`
	Nonce           int64  `json:"nonce"`
	BroadcastTxId   string `json:"broadcastTxId"`
}

type WalletPayment struct {
	Passphrase         string `json:"passphrase"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	Asset              string `json:"asset"`
	Quantity           uint64 `json:"quantity"`
	PaymentId          string `json:"paymentId"`
	PaymentTag         string `json:"paymentTag"`
	RequestId          string `json:"requestId"`
	Nonce              int64  `json:"nonce"`
}

type ActivateAddressStruct struct {
	Address       string `json:"address"`
	Amount        int64  `json:"amount"`
	ActivationId  string `json:"activationId"`
	BroadcastTxId string `json:"broadcastTxId"`
	Status        string `json:"status"`
	ErrorMessage  string `json:"errorMessage"`
	RequestId     string `json:"requestId"`
	Nonce         int64  `json:"nonce"`
}

type RippleAsset struct {
	Currency string `json:"currency,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
}
