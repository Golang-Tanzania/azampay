package main

import (
	"Golang-Tanzania/GoAzam"
	"fmt"
)

func main() {

	// initialize
	var transactionTester GoAzam.APICONTEXT

	transactionTester.LoadKeys("config.json")
	transactionTester.GenerateSessionID()

	// example mobile checkout
	exampleMobileCheckout := make(map[string]string)

	exampleMobileCheckout["accountNumber"] = "0700000000"
	exampleMobileCheckout["amount"] = "2000"
	exampleMobileCheckout["currency"] = "TZS"
	exampleMobileCheckout["externalId"] = "123"
	exampleMobileCheckout["provider"] = "TIGO"

	fmt.Println(transactionTester.MobileCheckout(exampleMobileCheckout))

	// example bank checkout
	exampleBankCheckout := make(map[string]string)

	exampleBankCheckout["amount"] = "10000"
	exampleBankCheckout["currencyCode"] = "TZS"
	exampleBankCheckout["merchantAccountNumber"] = "123321"
	exampleBankCheckout["merchantMobileNumber"] = "0700000000"
	exampleBankCheckout["otp"] = "1234"
	exampleBankCheckout["provider"] = "CRDB"
	exampleBankCheckout["ReferenceID"] = "123"

	fmt.Println(transactionTester.BankCheckout(exampleBankCheckout))

	// example Callback

	exampleCallback := make(map[string]string)

	exampleCallback["msisdn"] = "0178334"
	exampleCallback["amount"] = "2000"
	exampleCallback["message"] = "testing callback"
	exampleCallback["utilityref"] = "1282-123"
	exampleCallback["operator"] = "Airtel"
	exampleCallback["reference"] = "123-123"
	exampleCallback["transactionstatus"] = "success"
	exampleCallback["submerchantAcc"] = "01723113"

	exampleCallbackURL := "" // You need to set a webhook or fill provided URL
	fmt.Println(transactionTester.Callback(exampleCallback, exampleCallbackURL))

	// example Payment Partner

	fmt.Println(transactionTester.PaymentPartners())

}
