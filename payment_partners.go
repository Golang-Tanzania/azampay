package azampay

import (
	"context"
	"fmt"
)

func (c *Client) PaymentPartners(ctx context.Context) (*[]PayPartnersResponse, error) {

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s%s", c.APIBase, PayPartnersEndPoint), nil)
	if err != nil {
		return &[]PayPartnersResponse{}, err
	}
	resp := &[]PayPartnersResponse{}

	err = c.SendWithAuth(req, resp, nil)

	return resp, err
}
