package azampay

func NewAzamPay(keys Credentials) *AzamPay {
	api := &AzamPay{
		appName:      keys.AppName,
		clientID:     keys.ClientId,
		clientSecret: keys.ClientSecret,
		token:        keys.Token,
	}

	return api
}
