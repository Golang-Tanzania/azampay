package GoAzam

func (api *APICONTEXT) MobileCheckout(data map[string]string) string {
	return api.sendRequest(api.BaseURL, "/azampay/mno/checkout", data)
}

func (api *APICONTEXT) BankCheckout(data map[string]string) string {
	return api.sendRequest(api.BaseURL, "/azampay/bank/checkout", data)
}

func (api *APICONTEXT) Callback(data map[string]string, url string) string {
	return api.sendRequest(url, "/api/v1/Checkout/Callback", data)
}

func (api *APICONTEXT) PaymentPartners() string {
	return api.getRequest(api.BaseURL, "/api/v1/Partner/GetPaymentPartners")
}

/*
// This needs its own handling

func (api *APICONTEXT) PostCheckout(data map[string]interface{}) string {
	return api.sendRequest(api.BaseURL, "/api/v1/Partner/PostCheckout", data)
}
*/
