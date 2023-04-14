package azampay

import (
	"encoding/json"
	"errors"
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
