package GoAzam

import (
	"context"
	"fmt"
)

func (c *Client) NameLookUp(ctx context.Context, payload NameLookupPayload) (*NameLookupResponse, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, NameLookupEndPoint), payload)
	if err != nil {
		return &NameLookupResponse{}, err
	}
	resp := &NameLookupResponse{}

	err = c.SendWithAuth(req, resp, nil)

	return resp, err
}
