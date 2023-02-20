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

	// test mobile checkout
	testMobileCheckout := make(map[string]string)

	testMobileCheckout["accountNumber"] = "0700000000"
	testMobileCheckout["amount"] = "2000"
	testMobileCheckout["currency"] = "TZS"
	testMobileCheckout["externalId"] = "123"
	testMobileCheckout["provider"] = "TIGO"

	fmt.Println(transactionTester.MobileCheckout(testMobileCheckout))

	// test bank checkout
	testBankCheckout := make(map[string]string)

	testBankCheckout["amount"] = "10000"
	testBankCheckout["currencyCode"] = "TZS"
	testBankCheckout["merchantAccountNumber"] = "123321"
	testBankCheckout["merchantMobileNumber"] = "0700000000"
	testBankCheckout["otp"] = "1234"
	testBankCheckout["provider"] = "CRDB"
	testBankCheckout["ReferenceID"] = "123"

	fmt.Println(transactionTester.BankCheckout(testBankCheckout))

	// test Callback

	testCallback := make(map[string]string)

	testCallback["msisdn"] = "0178334"
	testCallback["amount"] = "2000"
	testCallback["message"] = "testing callback"
	testCallback["utilityref"] = "1282-123"
	testCallback["operator"] = "Airtel"
	testCallback["reference"] = "123-123"
	testCallback["transactionstatus"] = "success"
	testCallback["submerchantAcc"] = "01723113"

	testCallbackURL := "" // You need to set a webhook or fill provided URL
	fmt.Println(transactionTester.Callback(testCallback, testCallbackURL))

	// test Payment Partner

	fmt.Println(transactionTester.PaymentPartners())

}
