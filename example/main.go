package main

import (
	"Golang-Tanzania/GoAzam"
	"fmt"
)

func main() {
	var transactionTester GoAzam.APICONTEXT

	transactionTester.LoadKeys("config.json")
	transactionTester.GenerateSessionID()

	testMobileCheckout := make(map[string]string)

	testMobileCheckout["accountNumber"] = "0700000000"
	testMobileCheckout["amount"] = "2000"
	testMobileCheckout["currency"] = "TZS"
	testMobileCheckout["externalId"] = "123"
	testMobileCheckout["provider"] = "TIGO"

	fmt.Println(transactionTester.MobileCheckout(testMobileCheckout))

	testBankCheckout := make(map[string]string)

	testBankCheckout["amount"] = "10000"
	testBankCheckout["currencyCode"] = "TZS"
	testBankCheckout["merchantAccountNumber"] = "123321"
	testBankCheckout["merchantMobileNumber"] = "0700000000"
	testBankCheckout["otp"] = "1234"
	testBankCheckout["provider"] = "CRDB"
	testBankCheckout["ReferenceID"] = "123"

	fmt.Println(transactionTester.BankCheckout(testBankCheckout))
}
