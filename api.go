package azampay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
	Buffer       int `json:"buffer"`
	Debug        bool
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
		Buffer:       100,
	}

	return api
}

// MobileCheckout Function to send data to the MNO endpoint. It accepts a value of type
// MNOPayload and returns a value of type MNOResponse and an error if any.
func (api *AzamPay) MobileCheckout(payload MNOPayload) (*MNOResponse, error) {
	return Request[MNOResponse](api, &payload)
}

func (api *AzamPay) BankCheckout(payload BankCheckoutPayload) (*BankCheckoutResponse, error) {
	return Request[BankCheckoutResponse](api, &payload)
}

func (api *AzamPay) NameLookup(payload NameLookupPayload) (*NameLookupResponse, error) {
	return Request[NameLookupResponse](api, &payload)
}

func (api *AzamPay) CreateTransfer(payload CreateTransferPayload) (*CreateTransferResponse, error) {
	return Request[CreateTransferResponse](api, &payload)
}

type Params interface {
	data() interface{}
	endpoint() string
}

// Request TODO https://github.com/golang/go/issues/49085
// Todo redo this with better alternative
func Request[T any](api *AzamPay, payload Params) (*T, error) {
	//v := reflect.ValueOf(payload)

	//for i := 0; i < v.NumField(); i++ {
	//	if v.Field(i).String() == "" {
	//		return nil, fmt.Errorf("(Bank Checkout) Error: Field '%v' is required.", v.Type().Field(i).Name)
	//	}
	//}

	jsonParameters, err := json.Marshal(payload.data())

	if err != nil {
		return nil, err
	}

	url := api.BaseURL + payload.endpoint()

	if api.Debug {
		fmt.Printf("endpoint: %s\n", payload.endpoint())
		fmt.Printf("data: %s\n", string(jsonParameters))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParameters))
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %v", api.Bearer)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("X-API-KEY", api.token)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response *T

	if resp.StatusCode == 200 {
		decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&response)
		if decodeErr != nil {
			if decodeErr == io.EOF {
				return nil, fmt.Errorf("(Bank Checkout) Error: Server returned an empty body.")
			}
			return nil, decodeErr
		}

		if api.Debug {
			fmt.Printf("response: %+v\n", response)
		}

		return response, nil

	} else if resp.StatusCode == 400 {
		var badRequest *BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return nil, fmt.Errorf("(Bank Checkout) Error decoding badrequest: %w", err)
		}

		return nil, fmt.Errorf(badRequest.Error())
	} else if resp.StatusCode == 417 {
		var unauthorized *Unauthorized

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&unauthorized); err != nil {
			return nil, fmt.Errorf("(Bank Checkout) Error decoding unauthorized err: %w", err)
		}

		return nil, fmt.Errorf(unauthorized.Error())
	} else if resp.StatusCode == 500 {
		return nil, fmt.Errorf("(Bank Checkout) Internal Server Error: status code 500")
	} else {
		return nil, fmt.Errorf("(Bank Checkout) Error: status code %d", resp.StatusCode)
	}
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Clear discards all unprocessed incoming updates.
func (ch UpdatesChannel) Clear() {
	for len(ch) != 0 {
		<-ch
	}
}

// ListenForWebhook registers a http handler for a webhook.
func (api *AzamPay) ListenForWebhook(pattern string) UpdatesChannel {
	ch := make(chan Update, api.Buffer)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		update, err := api.HandleUpdate(r)

		if err != nil {
			errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(errMsg)
			return
		}

		ch <- *update
	})

	return ch
}

// ListenForWebhookRespReqFormat registers a http handler for a single incoming webhook.
func (api *AzamPay) ListenForWebhookRespReqFormat(w http.ResponseWriter, r *http.Request) UpdatesChannel {
	ch := make(chan Update, api.Buffer)

	func(w http.ResponseWriter, r *http.Request) {
		defer close(ch)

		update, err := api.HandleUpdate(r)
		if err != nil {
			errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(errMsg)
			return
		}

		ch <- *update
	}(w, r)

	return ch
}

// HandleUpdate parses and returns update received via webhook
func (api *AzamPay) HandleUpdate(r *http.Request) (*Update, error) {
	if r.Method != http.MethodPost {
		err := errors.New("wrong HTTP method required POST")
		return nil, err
	}

	var update Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return nil, err
	}

	return &update, nil
}
