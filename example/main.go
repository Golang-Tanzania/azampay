package main

import (
	"fmt"
	"os"

	"github.com/kateile/go-azampay"
)

func main() {
	// initialize
	transactionTester := azampay.NewAzamPay(false, azampay.Credentials{
		AppName:      os.Getenv("AZAM_APP_NAME"),
		ClientId:     os.Getenv("AZAM_CLIENT_ID"),
		ClientSecret: os.Getenv("AZAM_SECRET"),
		Token:        os.Getenv("AZAM_TOKEN"),
	})

	if err := transactionTester.GenerateSession(); err != nil {
		fmt.Println(err)
		return
	}

	// example mobile checkout
	var exampleMobileCheckout azampay.MNOPayload

	exampleMobileCheckout.AccountNumber = "0700000000"
	exampleMobileCheckout.Amount = "2000"
	exampleMobileCheckout.Currency = "TZS"
	exampleMobileCheckout.ExternalID = "123"
	exampleMobileCheckout.Provider = "TIGO"

	// The additional properties field are optional
	exampleAdditionalProperties := struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}{
		Name: "Sammy",
		Type: "Shark",
	}

	exampleMobileCheckout.AdditionalProperties = exampleAdditionalProperties

	mnoResult, err := transactionTester.MobileCheckout(exampleMobileCheckout)

	if err != nil {
		panic(err)
	}

	fmt.Println(mnoResult.Success)
	fmt.Println(mnoResult.Message)
	fmt.Println(mnoResult.TransactionID)

	// example bank checkout
	var exampleBankCheckout azampay.BankCheckoutPayload

	exampleBankCheckout.Amount = "10000"
	exampleBankCheckout.CurrencyCode = "TZS"
	exampleBankCheckout.MerchantAccountNumber = "123321"
	exampleBankCheckout.MerchantMobileNumber = "0700000000"
	exampleBankCheckout.MerchantName = "somebody"
	exampleBankCheckout.OTP = "1234"
	exampleBankCheckout.Provider = "CRDB"
	exampleBankCheckout.ReferenceID = "123"

	bankResult, err := transactionTester.BankCheckout(exampleBankCheckout)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bankResult.Success)
	fmt.Println(bankResult.Message)
	fmt.Println(bankResult.Data.Properties.ReferenceID)

	// example Callback
	var exampleCallback azampay.CallbackPayload

	exampleCallback.MSISDN = "0178334"
	exampleCallback.Amount = "2000"
	exampleCallback.Message = "testing callback"
	exampleCallback.UtilityRef = "1282-123"
	exampleCallback.Operator = "Airtel"
	exampleCallback.Reference = "123-123"
	exampleCallback.TransactionStatus = "success"
	exampleCallback.SubmerchantAcc = "01723113"

	// This domain should be the absolute path to your callback URL.
	// You can use the example server in this repository to test this endpoint.
	url := "http://localhost:8000/api/v1/Checkout/Callback"
	callbackResult, err := transactionTester.Callback(exampleCallback, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(callbackResult.Success)

	// example get Payment Partners

	examplePaymentPartners, err := transactionTester.PaymentPartners()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, paymentpartner := range examplePaymentPartners {
		fmt.Println(paymentpartner.PartnerName)
	}

	// example Post checkout

	var examplePostCheckout azampay.PostCheckoutPayload

	examplePostCheckout.AppName = "example"
	examplePostCheckout.Amount = "10000"
	examplePostCheckout.ClientID = "1234"
	examplePostCheckout.Currency = "TZS"
	examplePostCheckout.ExternalID = "30characterslong"
	examplePostCheckout.Language = "SW"
	examplePostCheckout.RedirectFailURL = "yoururl"
	examplePostCheckout.RedirectSuccessURL = "yoururl"
	examplePostCheckout.RequestOrigin = "yourorigin"
	examplePostCheckout.VendorName = "VendorName"
	examplePostCheckout.VendorID = "e9b57fab-1850-44d4-8499-71fd15c845a0"

	// Need to make list of shopping items if any
	shoppingList := []azampay.Item{
		{Name: "Mandazi"},
		{Name: "Sambusa"},
		{Name: "Mkate"},
	}
	examplePostCheckout.Cart.Items = append(examplePostCheckout.Cart.Items, shoppingList...)

	postCheckoutURL, err := transactionTester.PostCheckout(examplePostCheckout)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(postCheckoutURL)
}
