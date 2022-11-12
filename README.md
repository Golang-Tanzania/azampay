# Go Azam

A golang wrapper for Azam Payment Gateway. Still a work in progress.

```
package main

import (
	"fmt"
)

func main() {

	// Mobile Checkout Example

	var test1 APICONTEXT
	test1.LoadKeys("config.json")
	test1.generateSessionID()

	testMobile := make(map[string]string)

	testMobile["accountNumber"] = "0700000000"
	testMobile["amount"] = "2000"
	testMobile["currency"] = "TZS"
	testMobile["externalId"] = "123"
	testMobile["provider"] = "TIGO"

	fmt.Println(test1.MobileCheckout(testMobile))

	// Bank Checkout Example

	var test2 APICONTEXT
	test2.LoadKeys("config.json")
	test2.generateSessionID()

	testBank := make(map[string]string)

	testBank["amount"] = "2000"
	testBank["currencyCode"] = "TZS"
	testBank["merchantAccountNumber"] = "123321"
	testBank["merchantMobileNumber"] = "0700123123"
	testBank["otp"] = "1234"
	testBank["provider"] = "NMB"
	testBank["referenceId"] = "123"

	fmt.Println(test2.BankCheckout(testBank))

	//  Callback Example

	var test3 APICONTEXT
	test3.LoadKeys("config.json")
	test3.generateSessionID()

	testCallback := make(map[string]string)

	testCallback["msisdn"] = "0178823"
	testCallback["amount"] = "2000"
	testCallback["message"] = "any message"
	testCallback["utilityref"] = "1292-123"
	testCallback["operator"] = "Tigo"
	testCallback["reference"] = "123-123"
	testCallback["transactionstatus"] = "success"
	testCallback["submerchantAcc"] = "01723113"

	fmt.Println(test3.Callback(testCallback))

	// Payment Partner Example

	var test4 APICONTEXT
	test4.LoadKeys("config.json")
	test4.generateSessionID()

	testPaymentPartner := make(map[string]string)

	testPaymentPartner["currency"] = "TZS"
	testPaymentPartner["id"] = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	testPaymentPartner["logoUrl"] = "https://source.unsplash.com/random/200x200?sig=1"
	testPaymentPartner["partnerName"] = "Azampesa"
	testPaymentPartner["paymentPartnerId"] = "12031"
	testPaymentPartner["paymentVendorId"] = "123"
	testPaymentPartner["provider"] = "4"
	testPaymentPartner["status"] = "active"
	testPaymentPartner["vendorName"] = "AzamPesa"

	fmt.Println(test4.PaymentPartners(testPaymentPartner))
}
```
