package GoAzam

import (
	"net/http"
	"sync"
	"time"
)

const (
	APIBase              = "https://sandbox.azampay.co.tz"
	AuthenicateUrl       = "https://authenticator-sandbox.azampay.co.tz/AppRegistration/GenerateToken"
	MnoCheckoutEndPoint  = "/azampay/mno/checkout"
	BankCheckoutEndPoint = "/azampay/bank/checkout"
	PayPartnersEndPoint  = "/api/v1/Partner/GetPaymentPartners"
	PostCheckOutEndPoint = "/api/v1/Partner/PostCheckout"
)

type (
	TokenData struct {
		AccessToken string    `json:"accessToken"`
		Expire      time.Time `json:"expire"`
	}

	TokenRequest struct {
		AppName      string `json:"appName"`
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}
	TokenResponse struct {
		Data       TokenData `json:"data"`
		Message    string    `json:"message"`
		Success    bool      `json:"success"`
		StatusCode int       `json:"statusCode"`
	}

	ErrTokenResponse struct {
		Data       interface{} `json:"data"`
		Message    string      `json:"message"`
		Success    bool        `json:"success"`
		StatusCode int         `json:"statusCode"`
	}

	ErrorsMno struct {
		Amount        []interface{} `json:"amount"`
		ExternalID    []interface{} `json:"externalId"`
		AccountNumber []interface{} `json:"accountNumber"`
		Provider      []interface{} `json:"provider"`
	}

	ErrCheckOutResponse struct {
		ErrorsMno `json:"errors"`
		Type      string `json:"type"`
		TraceID   string `json:"traceId"`
		Title     string `json:"title"`
		Status    int    `json:"status"`
	}

	Client struct {
		// sync.Mutex
		mu           sync.Mutex
		Client       *http.Client
		AppName      string
		ClientID     string
		ClientSecret string
		APIBase      string
		Token        *TokenResponse
		TokenKey     string
	}

	MnoPayload struct {
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
		AdditionalProperties AdditionalProperties `json:"additionalProperties"`
	}

	// Data received from the server after a valid transaction
	MnoResponse struct {
		// Will be true is successful
		Success bool `json:"success"`
		// Each successful transaction will be given a valid transaction id. Can also be a string or null
		TransactionID string `json:"transactionId"`
		// This is the status message of checkout request. Can be a string or null
		Message string `json:"message"`
	}

	BankCheckoutPayload struct {
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
		AdditionalProperties AdditionalProperties `json:"additionalProperties"`
	}

	AdditionalProperties struct {
		Property1 any `json:"property1"`
		Property2 any `json:"property2"`
	}

	ReferenceID struct {
		// Reference ID of the transaction
		ReferenceID string `json:"ReferenceID"`
	}

	Properties struct {
		// List of properties
		Properties ReferenceID `json:"properties"`
	}

	// Data received from the server after a successful transaction
	BankCheckoutResponse struct {
		// will return true if successful
		Success bool `json:"success"`
		// message received from the server. Will be empty for sandbox
		Message string `json:"msg"`
		// data received from the server
		Data Properties `json:"data"`
	}

	// Payload that will be received from the payment partner payload.
	PayPartnersResponse struct {
		// ID of ther partner
		ID string `json:"id"`
		// Logo of the partner
		LogoURL string `json:"logoUrl"`
		// Name of the partner
		PartnerName string `json:"partnerName"`
		// Number of the provider
		Provider int64 `json:"provider"`
		// Name of the vendor
		VendorName string `json:"vendorName"`
		// ID of the payment vendor
		PaymentVendorID string `json:"paymentVendorId"`
		// ID of the payment partner
		PaymentPartnerID string `json:"paymentPartnerId"`
		// The callback url
		PaymentAcknowledgementRoute string `json:"paymentAcknowledgementRoute"`
		// Currency used
		Currency string `json:"currency"`
		// Status
		Status string `json:"status"`
		// Type of the vendor
		VendorType string `json:"vendorType"`
	}

	// Items in the shopping Cart
	Item struct {
		Name string `json:"name"`
	}

	// Shopping cart with multiple items
	Cart struct {
		// Items to be shopped
		Items []Item `json:"items"`
	}

	// Payload to be sent to the post checkout endpoint
	PostCheckoutPayload struct {
		// This is the amount that will be charged from the given account (required)
		Amount string `json:"amount"`
		// This is the application name (required)
		AppName string `json:"appName"`
		// Shopping cat with multiple items (required)
		Cart Cart `json:"cart"`
		// Unique identifier for the client (required)
		ClientID string `json:"clientId"`
		// Currency code that will convert amount into specific current (required)
		Currency string `json:"currency"`
		// 30 character long unique string (required)
		ExternalID string `json:"externalId"`
		// Language code to translate the application (required)
		Language string `json:"language"`
		// URL that be redirected to upon transaction failure (required)
		RedirectFailURL string `json:"redirectFailURL"`
		// URL to be directed to upon successful transaction (required)
		RedirectSuccessURL string `json:"redirectSuccessURL"`
		// URL which the request is being originated (required)
		RequestOrigin string `json:"requestOrigin"`
		// UUID to validate vendor (required)
		VendorID string `json:"vendorId"`
		// Name of the vendor (required)
		VendorName string `json:"vendorName"`
	}
)
