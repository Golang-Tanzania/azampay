package GoAzam

const (
	SandboxBaseURL    = "https://sandbox.azampay.co.tz"
	SandboxAuthURL    = "https://authenticator-sandbox.azampay.co.tz/AppRegistration/GenerateToken"
	ProductionBaseURL = "https://checkout.azampay.co.tz"
	ProductionAuthURL = "https://authenticator.azampay.co.tz/AppRegistration/GenerateToken"
)

type APICONTEXT struct {
	AppName      string
	ClientID     string
	ClientSecret string
	Token        string
	BaseURL      string
	Bearer       string
}

type kEYS struct {
	AppName      string
	ClientId     string
	ClientSecret string
	Token        string
}
