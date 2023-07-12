package azampay

import (
	"context"
	"fmt"
)

func (c *Client) TransactionalStatus(ctx context.Context, payload TransactionStatusQueries) (*[]TransactionStatusResponse, error) {

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, TransactionalStatusEndpoint), payload)
	if err != nil {
		return &[]TransactionStatusResponse{}, err
	}

	resp := &[]TransactionStatusResponse{}
	err = c.SendWithAuthQueriesParams(req, payload, resp)

	if err != nil {
		return &[]TransactionStatusResponse{}, err
	}

	return resp, err
}
