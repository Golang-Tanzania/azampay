package main

import (
	"fmt"
	"os"

	"github.com/kateile/go-azampay"
)

func main() {
	// initialize
	api := azampay.NewAzamPay(false, azampay.Credentials{
		AppName:      os.Getenv("AZAM_APP_NAME"),
		ClientId:     os.Getenv("AZAM_CLIENT_ID"),
		ClientSecret: os.Getenv("AZAM_SECRET"),
		Token:        os.Getenv("AZAM_TOKEN"),
	})

	api.Debug = true

	if err := api.GenerateSession(); err != nil {
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

	_, err := api.MobileCheckout(exampleMobileCheckout)

	if err != nil {
		panic(err)
	}

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

	_, err = api.BankCheckout(exampleBankCheckout)

	if err != nil {
		fmt.Println(err)
		return
	}

	// example get Payment Partners

	examplePaymentPartners, err := api.PaymentPartners()

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

	postCheckoutURL, err := api.PostCheckout(examplePostCheckout)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(postCheckoutURL)
}
