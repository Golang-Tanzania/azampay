package azampay

import "time"

// MNOPayload Payload to send to the MNO Checkout endpoint
type MNOPayload struct {
	// This is the account number/MSISDN that consumer will provide. The amount will be deducted from this account (required)
	AccountNumber string `json:"accountNumber"`
	// This is amount that will be charged from the given account (required)
	Amount string `json:"amount"`
	// This is the transaciton currency. Current support values are only TZS (required)
	Currency string `json:"currency"`
	// This id belongs to the calling application. Maximum Allowed length for this field is 128 ascii characters (required)
	ExternalID string `json:"externalId"`
	// Only providers available are Airtel, Tigo, Halopesa and Azampesa (required)
	Provider string `json:"provider"`
	// This is additional data you can provide (Optional)
	AdditionalProperties interface{} `json:"additionalProperties"`
}

func (p *MNOPayload) data() interface{} {
	return p
}

func (p *MNOPayload) endpoint() string {
	return "/azampay/mno/checkout"
}

// MNOResponse Data received from the server after a valid transaction
type MNOResponse struct {
	// Will be true is successful
	Success bool `json:"success"`
	// Each successful transaction will be given a valid transaction id. Can also be a string or null
	TransactionID string `json:"transactionId"`
	// This is the status message of checkout request. Can be a string or null
	Message string `json:"message"`
}

// BankCheckoutPayload Payload to send to the bank checkout endpoint
type BankCheckoutPayload struct {
	// This is amount that will be charged from the given account (required)
	Amount string `json:"amount"`

	// Code of currency (required)
	CurrencyCode string `json:"currencyCode"`

	// This is the account number/MSISDN that consumer will provide. The amount will be deducted from this account (required)
	MerchantAccountNumber string `json:"merchantAccountNumber"`

	// Mobile number (required)
	MerchantMobileNumber string `json:"merchantMobileNumber"`

	// The name of the customer (optional)
	MerchantName string `json:"merchantName"`

	// One time password (required)
	OTP string `json:"otp"`

	// Bank provider. Currently on CRDB and NMB are supported (required)
	Provider string `json:"provider"`

	// This id belongs to the calling application. Maximum Allowed length for this field is 128 ascii characters (Optional)
	ReferenceID string `json:"referenceId"`

	// This is additional data you can provide (Optional)
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
}

func (p *BankCheckoutPayload) data() interface{} {
	return p
}

func (p *BankCheckoutPayload) endpoint() string {
	return "/azampay/bank/checkout"
}

type ReferenceID struct {
	// Reference ID of the transaction
	ReferenceID string `json:"ReferenceID"`
}

type Properties struct {
	// List of properties
	Properties ReferenceID `json:"properties"`
}

// BankCheckoutResponse Data received from the server after a successful transaction
type BankCheckoutResponse struct {
	// will return true if successful
	Success bool `json:"success"`
	// message received from the server. Will be empty for sandbox
	Message string `json:"msg"`
	// data received from the server
	Data Properties `json:"data"`
}

// NameLookup The endpoint to lookup the name associated with
// a bank account or mobile money
type NameLookup struct {
	// Bank account number or mobile money number
	BankName string `json:"bankName"`
	// Bank name or mobile money name associated with the account
	AccountNumber string `json:"accountNumber"`
}

type Source struct {
	// Country code of the source country
	CountryCode string `json:"countryCode"`
	// Full name of the account holder
	FullName string `json:"fullName"`
	// The name of the bank where the source account is held.
	// Current options are 'tigo', 'airtel', 'azampesa'
	BankName string `json:"bankName"`
	// The account number of the source account
	AccountNumber string `json:"accountNumber"`
	// The currency in which the transfer is made
	Currency string `json:"currency"`
}

type Destination struct {
	// Country code of the destination account
	CountryCode string `json:"countryCode"`
	// The full name of the account holder
	FullName string `json:"fullName"`
	// The bank where the destination account is held
	// Current options are 'tigo', 'airtel', 'azampesa'
	BankName string `json:"bankName"`
	// The account number of the destination account
	AccountNumber string `json:"accountNumber"`
	// The currency in which the transfer is made
	Currency string `json:"currency"`
}

type TransferDetails struct {
	// The type of the transfer eg: SWIFT, SEPA etc
	Type string `json:"type"`
	// The amount to be transfered
	Amount int `json:"amount"`
	// The date when transfer is made
	Date time.Time `json:"date"`
}

// DisbursePayload Allows for transfer of money from other countries
// to Tanzania.
type DisbursePayload struct {
	// Contains information about the source account
	Source Source `json:"source"`
	// Contains information about the destination account
	Destination Destination `json:"destination"`
	// Contains information about the transfer
	TransferDetails TransferDetails `json:"transferDetails"`
	// An external reference ID to track the transaction
	ExternalReferenceID string `json:"externalReferenceId"`
	// Any Remarks to be included in the transaction
	Remarks string `json:"remarks"`
}

func (p *DisbursePayload) data() interface{} {
	return p
}

func (p *DisbursePayload) endpoint() string {
	return "/azampay/createtransfer"
}

type DisburseResponseObject struct {
	// A string containing the status of the transaction.
	Data string `json:"data"`
	// A string containing a human-readable message describing the response.
	Message string `json:"message"`
	// A boolean indicating whether the request was successful or not.
	Success bool `json:"success"`
	// A string containing a human-readable message describing the response.
	StatusCode string `json:"statusCode"`
}

type DisburseResponse = []DisburseResponseObject

// Update Payload to be sent to the callback endpoint
type Update struct {
	// This is amount that will be charged from the given account.
	Amount string `json:"amount"`
	// This is the transaction description message
	Message      string `json:"message"`
	MNOReference string `json:"mnoreference"`
	// This is the account number/MSISDN that consumer will provide. The amount will be deducted from this account
	MSISDN string `json:"msisdn"`
	// Only operators available are Airtel, Tigo, Halopesa and Azampesa
	Operator string `json:"operator"`
	// This is the transaction ID
	Reference string `json:"reference"`
	// This field is reserved for future use according to the Azampay documentation
	SubmerchantAcc string `json:"submerchantAcc"`
	// Whether the transaction was a success or fail
	TransactionStatus string `json:"transactionStatus"`
	// This is the ID that belongs to the calling application
	UtilityRef string `json:"utilityref"`
	// This is additional JSON data that calling application can provide. This is optional.
	Properties interface{} `json:"additionalProperties,omitempty"`
}
