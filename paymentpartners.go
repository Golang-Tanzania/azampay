package GoAzam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	PaymentPartner struct {
		ID                          string `json:"id"`
		LogoURL                     string `json:"logoUrl"`
		PartnerName                 string `json:"partnerName"`
		Provider                    int64  `json:"provider"`
		VendorName                  string `json:"vendorName"`
		PaymentVendorID             string `json:"paymentVendorId"`
		PaymentPartnerID            string `json:"paymentPartnerId"`
		PaymentAcknowledgementRoute string `json:"paymentAcknowledgementRoute"`
		Currency                    string `json:"currency"`
		Status                      string `json:"status"`
		VendorType                  string `json:"vendorType"`
	}

	PaymentPartners []PaymentPartner
)

func (api *APICONTEXT) PaymentPartners() (PaymentPartners, error) {

	url := api.BaseURL + "/api/v1/Partner/GetPaymentPartners"

	req, err := http.NewRequest("GET", url, nil)

	bearer := fmt.Sprintf("Bearer %v", api.Bearer)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("X-API-KEY", api.Token)
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

	var paymentpartners PaymentPartners

	if resp.StatusCode == 200 {
		decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&paymentpartners)

		if decodeErr != nil {
			if decodeErr == io.EOF {
				return nil, fmt.Errorf("Error: Server returned an empty body")
			}
			return nil, decodeErr
		}

		return paymentpartners, nil

	} else if resp.StatusCode == 400 {
		var badRequest BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return nil, fmt.Errorf("Error decoding badrequest: %w", err)
		}

		return nil, fmt.Errorf(badRequest.Error())

	} else if resp.StatusCode == 417 {
		var unauthorized *Unauthorized

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&unauthorized); err != nil {
			return nil, fmt.Errorf("Error decoding unauthorized err: %w", err)
		}

		return nil, fmt.Errorf(unauthorized.Error())

	} else if resp.StatusCode == 500 {

		return nil, fmt.Errorf("Internal Server Error: status code 500")

	} else {

		return nil, fmt.Errorf("Error: status code %d", resp.StatusCode)

	}

}
