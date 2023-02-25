package GoAzam

// Base URLs
const (
	// Sandbox URLs

	SandboxBaseURL = "https://sandbox.azampay.co.tz"
	SandboxAuthURL = "https://authenticator-sandbox.azampay.co.tz/AppRegistration/GenerateToken"

	// Production URLs

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
	AppName      string
	ClientId     string
	ClientSecret string
	Token        string
}


// This will return provider when phonenumber is passed(07XX-XXX-XXX)
func (api *APICONTEXT) GetProvider(phoneNumber string) string {

	phoneCode := phoneNumber[:3]

	if phoneCode == "065" || phoneCode == "067" || phoneCode == "071" { //TIGO
		return "Tigo"
	} else if phoneCode == "074" || phoneCode == "075" || phoneCode == "076" { //Vodacom
		return "Mpesa"
	} else if phoneCode == "064" || phoneCode == "077" { //Zantel
		return "Tigo" //@@Since Zantel and Tigo are one Company
	} else if phoneCode == "062" || phoneCode == "061" { //Halopesa
		return "Halopesa"
	} else if phoneCode == "068" || phoneCode == "069" || phoneCode == "078" { //Airtel
		return "Airtel"
	} else {
		return ""
	}
}
