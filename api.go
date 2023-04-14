package azampay

// Base URLs
const (
	// SandboxBaseURL Sandbox URLs
	SandboxBaseURL = "https://sandbox.azampay.co.tz"
	SandboxAuthURL = "https://authenticator-sandbox.azampay.co.tz/AppRegistration/GenerateToken"

	// ProductionBaseURL Production URLs
	ProductionBaseURL = "https://checkout.azampay.co.tz"
	ProductionAuthURL = "https://authenticator.azampay.co.tz/AppRegistration/GenerateToken"
)

// AzamPay This will be the API type to initialize
// config variables and hold the bearer token
type AzamPay struct {
	appName      string
	clientID     string
	clientSecret string
	token        string
	BaseURL      string
	Bearer       string
	Expiry       string
	IsLive       bool
}

// Credentials A helper struct to read values from the
type Credentials struct {
	AppName      string
	ClientId     string
	ClientSecret string
	Token        string
}

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
