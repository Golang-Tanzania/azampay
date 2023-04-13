package azampay

// NameLookup The endpoint to lookup the name associated with
// a bank account or mobile money
type NameLookup struct {
	// Bank account number or mobile money number
	BankName string `json:"bankName"`
	// Bank name or mobile money name associated with the account
	AccountNumber string `json:"accountNumber"`
}
