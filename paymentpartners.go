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
		// ID of ther partner
		ID string `json:"id"`
		// Logo of the partner
		LogoURL string `json:"logoUrl"`
		// Name of the partner
		PartnerName string `json:"partnerName"`
		// Number of the provider
		Provider int64 `json:"provider"`
		// Name of the vendor
		VendorName string `json:"vendorName"`
		// ID of the payment vendor
		PaymentVendorID string `json:"paymentVendorId"`
		// ID of the payment partner
		PaymentPartnerID string `json:"paymentPartnerId"`
		// The callback url
		PaymentAcknowledgementRoute string `json:"paymentAcknowledgementRoute"`
		// Currency used
		Currency string `json:"currency"`
		// Status
		Status string `json:"status"`
		// Type of the vendor
		VendorType string `json:"vendorType"`
	}

	// List of the payment partners
	PaymentPartners []PaymentPartner
)

func (api *APICONTEXT) PaymentPartners() (PaymentPartners, error) {

	url := api.BaseURL + "/api/v1/Partner/GetPaymentPartners"

	req, err := http.NewRequest("GET", url, nil)

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
