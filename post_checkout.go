package GoAzam

import (
	"context"
	"fmt"
)

func (c *Client) PostCheckout(ctx context.Context, payload PostCheckoutPayload) (string, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, PostCheckOutEndPoint), payload)
	if err != nil {
		return "", err
	}

	data, err := c.SendWithAuthMultiReturns(req)

	if err != nil {
		return "", err
	}

	return data, err
}
