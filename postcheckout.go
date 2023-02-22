package GoAzam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	// Shopping cart with multiple items
	Cart struct {
		// Items to be shopped
		Items map[string]string `json:"items"`
	}

	// Payload to be sent to the post checkout endpoint
	PostCheckoutPayload struct {
		// This is the amount that will be charged from the given account (required)
		Amount string `json:"amount"`
		// This is the application name (required)
		AppName string `json:"appName"`
		// Shopping cat with multiple items (required)
		Cart Cart `json:"cart"`
		// Unique identifier for the client (required)
		ClientID string `json:"clientId"`
		// Currency code that will convert amount into specific current (required)
		Currency string `json:"currency"`
		// 30 character long unique string (required)
		ExternalID string `json:"externalId"`
		// Language code to translate the application (required)
		Language string `json:"language"`
		// URL that be redirected to upon transaction failure (required)
		RedirectFailURL string `json:"redirectFailURL"`
		// URL to be directed to upon successful transaction (required)
		RedirectSuccessURL string `json:"redirectSuccessURL"`
		// URL which the request is being originated (required)
		RequestOrigin string `json:"requestOrigin"`
		// UUID to validate vendor (required)
		VendorID string `json:"vendorId"`
		// Name of the vendor (required)
		VendorName string `json:"vendorName"`
	}
)

// PostCheckout to function to get the post checkout URL.
// It accepts a payload of type PostCheckoutPayload and
// returns the checkout url as a string and an error if any.
func (api *APICONTEXT) PostCheckout(postcheckoutpayload PostCheckoutPayload) (string, error) {

	jsonParameters, err := json.Marshal(postcheckoutpayload)

	if err != nil {
		return "", err
	}

	url := api.BaseURL + "/api/v1/Partner/PostCheckout"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonParameters)))

	bearer := fmt.Sprintf("Bearer %v", api.Bearer)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("X-API-KEY", api.token)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var postcheckouturl string

	if resp.StatusCode == 200 {
		postcheckouturl = string(body)
		return postcheckouturl, nil

	} else if resp.StatusCode == 400 {
		var badRequest *BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return "", fmt.Errorf("(Post Checkout) Error decoding badrequest: %w", err)
		}

		return "", fmt.Errorf(badRequest.Error())

	} else if resp.StatusCode == 417 {
		var unauthorized *Unauthorized

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&unauthorized); err != nil {
			return "", fmt.Errorf("(Post Checkout) Error decoding unauthorized err: %w", err)
		}

		return "", fmt.Errorf(unauthorized.Error())

	} else if resp.StatusCode == 500 {

		return "", fmt.Errorf("(Post Checkout) Internal Server Error: status code 500")

	} else {

		return "", fmt.Errorf("(Post Checkout) Error: status code %d", resp.StatusCode)

	}

}
