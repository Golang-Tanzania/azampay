package azampay

import (
	"context"
	"fmt"
)

func (c *Client) Disburse(ctx context.Context, payload DisbursePayload) (*DisburseResponse, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, DisburseEndpoint), payload)
	if err != nil {
		return &DisburseResponse{}, err
	}
	resp := &DisburseResponse{}

	err = c.SendWithAuth(req, resp, nil)

	return resp, err
}
