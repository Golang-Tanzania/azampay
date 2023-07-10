package GoAzam

import (
	"context"
	"fmt"
)

func (c *Client) BankCheckout(ctx context.Context, payload BankCheckoutPayload) (*BankCheckoutResponse, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, BankCheckoutEndPoint), payload)
	if err != nil {
		return &BankCheckoutResponse{}, err
	}
	resp := &BankCheckoutResponse{}

	errCheckOutResponse := &ErrCheckOutResponse{}

	err = c.SendWithAuth(req, resp, errCheckOutResponse)

	// Just check one of the field if its not empty(only one field )
	if errCheckOutResponse.Title != "" {
		return &BankCheckoutResponse{}, fmt.Errorf("success %t , status code %d, %s, %+v ,%+v", resp.Success, errCheckOutResponse.Status, errCheckOutResponse.Title, errCheckOutResponse.ErrorsMno, errCheckOutResponse.TraceID)
	}

	return resp, err
}
