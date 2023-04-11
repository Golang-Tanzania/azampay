package azampay

// This API allows you to retrieve the status of a disbursement transaction made through Azampay
type TransactionStatus struct {
	// The name of the mobile network operator (MNO) used
	// to make the disbursement request
	BankName string `json:"bankName"`
	// The transaction ID you received when making the
	// disbursement request
	PgReferenceID string `json:"pgReferenceId"`
}
