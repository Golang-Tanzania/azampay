package GoAzam

func NewAPI(keys AzamCredentials) *APICONTEXT {
	api := &APICONTEXT{
		appName:      keys.AppName,
		clientID:     keys.ClientId,
		clientSecret: keys.ClientSecret,
		token:        keys.Token,
	}

	return api
}
