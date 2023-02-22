package GoAzam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	MNOPayload struct {
		// This is the account number/MSISDN that consumer will provide. The amount will be deducted from this account (required)
		AccountNumber string `json:"accountNumber"`
		// This is amount that will be charged from the given account (required)
		Amount string `json:"amount"`
		// This is the transaciton currency. Current support values are only TZS (required)
		Currency string `json:"currency"`
		// This id belongs to the calling application. Maximum Allowed length for this field is 128 ascii characters (required)
		ExternalID string `json:"externalId"`
		// Only providers available are Airtel, Tigo, Halopesa and Azampesa (required)
		Provider string `json:"provider"`
	}

	MNOResponse struct {
		// Will be true is successful
		Success bool `json:"success"`
		// Each successful transaction will be given a valid transaction id. Can also be a string or null
		TransactionID string `json:"transactionId"`
		// This is the status message of checkout request. Can be a string or null
		Message string `json:"message"`
	}
)

func (api *APICONTEXT) MobileCheckout(mnopayload MNOPayload) (*MNOResponse, error) {

	jsonParameters, err := json.Marshal(mnopayload)

	if err != nil {
		return nil, err
	}

	url := api.BaseURL + "/azampay/mno/checkout"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonParameters)))

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

	var mnoresponse *MNOResponse

	if resp.StatusCode == 200 {
		decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&mnoresponse)

		if decodeErr != nil {
			if decodeErr == io.EOF {
				return nil, fmt.Errorf("(MNO) Error: Server returned an empty body")
			}
			return nil, decodeErr
		}

		return mnoresponse, nil

	} else if resp.StatusCode == 400 {
		var badRequest *BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return nil, fmt.Errorf("(MNO) Error decoding badrequest: %w", err)
		}

		return nil, fmt.Errorf(badRequest.Error())

	} else if resp.StatusCode == 417 {
		var unauthorized *Unauthorized

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&unauthorized); err != nil {
			return nil, fmt.Errorf("(MNO) Error decoding unauthorized err: %w", err)
		}

		return nil, fmt.Errorf(unauthorized.Error())

	} else if resp.StatusCode == 500 {

		return nil, fmt.Errorf("(MNO) Internal Server Error: status code 500")

	} else {

		return nil, fmt.Errorf("(MNO) Error: status code %d", resp.StatusCode)

	}

}
