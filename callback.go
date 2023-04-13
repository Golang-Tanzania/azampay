package azampay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	// CallbackPayload Payload to be sent to the callback endpoint
	CallbackPayload struct {
		// This is amount that will be charged from the given account.
		Amount string `json:"amount"`
		// This is the transaction description message
		Message string `json:"message"`
		// This is the account number/MSISDN that consumer will provide. The amount will be deducted from this account
		MSISDN string `json:"msisdn"`
		// Only operators available are Airtel, Tigo, Halopesa and Azampesa
		Operator string `json:"operator"`
		// This is the transaction ID
		Reference string `json:"reference"`
		// This field is reserved for future use according to the Azampay documentation
		SubmerchantAcc string `json:"submerchantAcc"`
		// Whether the transaction was a success or fail
		TransactionStatus string `json:"transactionStatus"`
		// This is the ID that belongs to the calling application
		UtilityRef string `json:"utilityref"`
	}

	// CallbackResponse There is no documentation on how the response is shaped, but I am assuming there will be a success field
	CallbackResponse struct {
		// Will be true is successful
		Success bool `json:"success"`
	}
)

// Callback Function to access the callback endpoint. It accepts a parameter of type Callback,
// an absolute URL to the checkout endpoint and will return a value of type CallbackResponse
// and an error if any
func (api *AzamPay) Callback(callbackpayload CallbackPayload, url string) (*CallbackResponse, error) {

	jsonParameters, err := json.Marshal(callbackpayload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParameters))

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

	var callbackresponse *CallbackResponse

	if resp.StatusCode == 200 {
		decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&callbackresponse)

		if decodeErr != nil {
			if decodeErr == io.EOF {
				return nil, fmt.Errorf("(Callback) Error: Server returned an empty body")
			}
			return nil, decodeErr
		}

		return callbackresponse, nil

	} else if resp.StatusCode == 400 {
		var badRequest *BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return nil, fmt.Errorf("(Callback) Error decoding badrequest: %w", err)
		}

		return nil, fmt.Errorf(badRequest.Error())

	} else if resp.StatusCode == 417 {
		var unauthorized *Unauthorized

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&unauthorized); err != nil {
			return nil, fmt.Errorf("(Callback) Error decoding unauthorized err: %w", err)
		}

		return nil, fmt.Errorf(unauthorized.Error())

	} else if resp.StatusCode == 500 {

		return nil, fmt.Errorf("(Callback) Internal Server Error: status code 500")

	} else {

		return nil, fmt.Errorf("(Callback) Error: status code %d", resp.StatusCode)

	}

}
