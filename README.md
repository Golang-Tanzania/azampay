<h1 align="center">AZAMPAY</h1>

<p align="center">
<a href="https://github.com/Golang-Tanzania/azampay"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg"></a>
<a href="https://github.com/Golang-Tanzania/azampay"><img src="https://img.shields.io/github/go-mod/go-version/gomods/athens.svg"></a>
<a href="https://pkg.go.dev/github.com/Golang-Tanzania/azampay"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a><br><br>
A golang client to access the Azam Payment Gateway. Made with love for gophers ❤️.
</p>

## Introduction

[AzamPay](https://developerdocs.azampay.co.tz/redoc) is specialized in the development of end-to-end online payment management solutions for companies operating in East Africa. They provide an API which allows developers to integrate their system's to Azampay's gateway.

This is a Golang client which significantly simplifies access and integration to Azampay's gateway.

<p align="center">
<img src="./assets/azampay-api-flow.svg">
</p>

## Features of azampay

- Make mobile network checkouts
- Make bank checkouts 
- Manage callback URLs after a transaction is confirmed
- Return a list of registered partners of the provided merchant 
- Create post checkout URLs for payments
- Transfer of money from other countries to Tanzania
- Lookup the name associated with a bank account or Mobile Money account
- Retrieve the status of a disbursement transaction made through AzamPay

<img src="./assets/azampay.gif" align="center">

## Pre-Requisites

- Sign up for a developer account with [Azampay](https://developers.azampay.co.tz/)
- Register an app to get credentials
- Use the provided credentials to access the API. 

## Installation

*Note: Ensure you have initialized your go code with `go mod init`*

Install the package with the `go get` command as shown below:
```sh 
go get github.com/Golang-Tanzania/azampay@latest
```

Then import it as follows:
```go
package main 

import (
    "github.com/Golang-Tanzania/azampay"
)
```

## Authentication

### Token Generation

```go
appName := "Your app name from azamm pay"
clientId := "Client id from azam pay"
clientSecret := "Client secret from azam pay"
tokenKey := "Your token from azam pay"
client, err := azampay.NewClient(appName, clientId, clientSecret, tokenKey)

if err != nil {
		panic(err)
}

ctx := context.Background()
_, err = client.GetAccessToken(ctx)

if err != nil {
		fmt.Println(err)
	}

```

## Transactions 

### MNO Checkout 


```go
// example MNO checkout
	exampleMNOCheckout := azampay.MnoPayload{
		AccountNumber: "0654001122",
		Amount:        "2000",
		Currency:      "TZS",
		ExternalID:    "123",
		Provider:      "Tigo",
	}

	res, err := client.MnoCheckout(ctx, exampleMNOCheckout)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

```

### Bank Checkout

```go
// example Bank checkout
exampleBankCheckout := azampay.BankCheckoutPayload{
		Amount:                "10000",
		CurrencyCode:          "TZS",
		MerchantAccountNumber: "123321",
		MerchantMobileNumber:  "0700000000",
		MerchantName:          "somebody",
		OTP:                   "1234",
		Provider:              "CRDB",
		ReferenceID:           "123",
	}

	res, err := client.BankCheckout(ctx, exampleBankCheckout)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

```


### Payment Partners 

```go

// example get registered partners of the provided merchant 
rp, err := client.PaymentPartners(ctx)

if err != nil {
	fmt.Println(err)
}

fmt.Println(res)

```

### Post Checkout 

```go
// example Post Checkout 
payload := azampay.PostCheckoutPayload{
		AppName:            "example",
		Amount:             "10000",
		ClientID:           "1234",
		Currency:           "TZS",
		ExternalID:         "30characterslong",
		Language:           "SW",
		RedirectFailURL:    "yoururl",
		RedirectSuccessURL: "yourrul",
		RequestOrigin:      "yourorigin",
		VendorName:         "VendorName",
		VendorID:           "e9b57fab-1850-44d4-8499-71fd15c845a0",
	}

	shoppingList := []azampay.Item{
		{Name: "Mandazi"},
		{Name: "Sambusa"},
		{Name: "Mkate"},
	}
    payload.Cart.Items = shoppingList
	
	res ,err := client.PostCheckout(ctx,payload) 

	if err != nil {
		fmt.Println(res)
	}

	fmt.Println(res)

```

### Disburse
```go 

//example Transfer of money from other countries to Tanzania
disbursePayload := azampay.DisbursePayload{
		Source: azampay.Source{
			CountryCode:   "US",
			FullName:      "John Doe",
			BankName:      "Bank of America",
			AccountNumber: "123456789",
			Currency:      "USD",
		},
		Destination: azampay.Destination{
			CountryCode:   "TZ",
			FullName:      "Jane Doe",
			BankName:      "Azania Bank",
			AccountNumber: "987654321",
			Currency:      "TZS",
		},
		TransferDetails: azampay.TransferDetails{
			Type:   "SWIFT",
			Amount: 5000,
			Date:   time.Date(2023, 7, 11, 0, 0, 0, 0, time.UTC),
		},
		ExternalReferenceID: "123",
		Remarks:             "Payment for goods",
	}



	res ,err := client.Disburse(ctx, disbursePayload)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(res)


```
### Name Lookup

```go

// example to lookup the name associated with a bank account or Mobile Money account.
res, err := client.NameLookUp(ctx, azampay.NameLookupPayload{
		AccountNumber: "0654000000",
		BankName:      "Tigo",
	})

	if err != nil {
       fmt.Println(err)

	}

fmt.Println(res)

```


### Get Transaction Status

```go

// example to retrieve the status of a disbursement transaction made through AzamPay.
queries := azampay.TransactionStatusQueries{

    BankName : "YOUR_MNO_NAME_HERE"
    PgReferenceID : "YOUR_TRANSACTION_ID_HERE"
}
res, err := client.TransactionalStatus(ctx,queries )

if err != nil {
	fmt.Println(err)
}

fmt.Println(res)

```



## Issues

If you notice any issues with the package kindly notify us as soon as possible.

## Credits

- [Hopertz](https://github.com/Hopertz)
- [Avicenna](https://github.com/AvicennaJr)
- All other [contributors](https://github.com/Golang-Tanzania/azampay/graphs/contributors)


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
