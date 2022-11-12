package GoAzam

func (api *APICONTEXT) MobileCheckout(data map[string]string) string {
	return api.sendRequest("/azampay/mno/checkout", data)
}

func (api *APICONTEXT) BankCheckout(data map[string]string) string {
	return api.sendRequest("/azampay/bank/checkout", data)
}

func (api *APICONTEXT) Callback(data map[string]string) string {
	return api.sendRequest("/api/v1/Checkout/Callback", data)
}

func (api *APICONTEXT) PaymentPartners(data map[string]string) string {
	return api.sendRequest("/api/v1/Partner/GetPaymentPartners", data)
}

/*
// This needs its own handling

func (api *APICONTEXT) PostCheckout(data map[string]interface{}) string {
	return api.sendRequest("/api/v1/Partner/PostCheckout", data)
}
*/
