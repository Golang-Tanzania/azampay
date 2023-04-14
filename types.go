package azampay

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
