package GoAzam

import (
	"context"
	"fmt"
)

func (c *Client) MnoCheckout(ctx context.Context, payload MnoPayload) (*MnoResponse, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, MnoCheckoutEndPoint), payload)
	if err != nil {
		return &MnoResponse{}, err
	}
	resp := &MnoResponse{}

	errCheckOutResponse := &ErrCheckOutResponse{}

	err = c.SendWithAuth(req, resp, errCheckOutResponse)

	// Just check one of the field if its not empty(only one field )
	if errCheckOutResponse.Title != "" {
		return &MnoResponse{}, fmt.Errorf("success %t , status code %d, %s, %+v ,%+v", resp.Success, errCheckOutResponse.Status, errCheckOutResponse.Title, errCheckOutResponse.ErrorsMno, errCheckOutResponse.TraceID)
	}

	return resp, err
}
