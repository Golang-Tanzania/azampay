package GoAzam

import "time"

type (
	// Allows for transfer of money from other countries
	// to Tanzania.
	DisbursePayload struct {
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
	Source struct {
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
	Destination struct {
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
	TransferDetails struct {
		// The type of the transfer eg: SWIFT, SEPA etc
		Type string `json:"type"`
		// The amount to be transfered
		Amount int `json:"amount"`
		// The date when transfer is made
		Date time.Time `json:"date"`
	}
)
