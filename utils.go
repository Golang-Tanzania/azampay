package GoAzam

const (
	SandboxBaseURL    = "https://sandbox.azampay.co.tz"
	SandboxAuthURL    = "https://authenticator-sandbox.azampay.co.tz/AppRegistration/GenerateToken"
	ProductionBaseURL = "https://checkout.azampay.co.tz"
	ProductionAuthURL = "https://authenticator.azampay.co.tz/AppRegistration/GenerateToken"
)

// This will be the API type to initialize
// config variables and hold the bearer token
type APICONTEXT struct {
	appName      string
	clientID     string
	clientSecret string
	token        string
	BaseURL      string
	Bearer       string
	Expiry       string
}

// A helper struct to read values from the
// config.json file
type keys struct {
	appName      string
	clientId     string
	clientSecret string
	token        string
}
