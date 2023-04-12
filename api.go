package azampay

func NewAzamPay(isLive bool, keys Credentials) *AzamPay {
	api := &AzamPay{
		appName:      keys.AppName,
		clientID:     keys.ClientId,
		clientSecret: keys.ClientSecret,
		token:        keys.Token,
		IsLive:       isLive,
	}

	return api
}
