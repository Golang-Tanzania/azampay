package azampay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// NewClient returns new Client struct
func NewClient(appName, clientID, clientSecret, tokenKey string) (*Client, error) {
	if appName == "" || clientID == "" || clientSecret == "" || tokenKey == "" {
		return nil, errors.New("appName , clientID, clientSecret, tokenKey are required to create a Client")
	}

	return &Client{
		Client:       &http.Client{},
		AppName:      appName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenKey:     tokenKey,
		APIBase:      APIBase,
	}, nil
}

func (c *Client) GetAccessToken(ctx context.Context) (*TokenResponse, error) {

	payload := &TokenRequest{
		AppName:      c.AppName,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	}

	req, err := c.NewRequest(ctx, "POST", AuthenicateUrl, payload)
	if err != nil {
		return &TokenResponse{}, err
	}

	response := &TokenResponse{}

	errResponse := &ErrTokenResponse{}

	err = c.Send(req, response, errResponse)

	// Set Token for current Client
	if response.Data.AccessToken != "" {
		c.Token = response
		return response, err

	}

	return &TokenResponse{}, fmt.Errorf("status code %d, %s", errResponse.StatusCode, errResponse.Message)

}

// SetHTTPClient sets *http.Client to current client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

// SendWithAuth makes a request to the API using clientID:secret basic auth
func (c *Client) SendWithAuth(req *http.Request, v interface{}, e interface{}) error {

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token.Data.AccessToken))
	req.Header.Set("X-API-Key", c.TokenKey)

	return c.Send(req, v, e)
}

func (c *Client) Send(req *http.Request, v interface{}, e interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode == 500 {
		return errors.New("internal Server Error")

	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		data, err = io.ReadAll(resp.Body)

		if e == nil {
			return fmt.Errorf("unknown error (%s), status code: %d", string(data), resp.StatusCode)
		}

		if err == nil && len(data) > 0 {
			err := json.Unmarshal(data, e)
			if err != nil {
				return err
			}
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}
	return json.NewDecoder(resp.Body).Decode(v)

}

// SendWithAuth makes a request to the API using clientID:secret basic auth
func (c *Client) SendWithAuthStringReturns(req *http.Request) (string, error) {

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token.Data.AccessToken))

	req.Header.Set("X-API-Key", c.TokenKey)

	return c.SendWithStringReturns(req)
}

func (c *Client) SendWithStringReturns(req *http.Request) (string, error) {
	var (
		err  error
		resp *http.Response
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode == 500 {
		return "", errors.New("internal Server Error")

	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		return "", fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *Client) SendWithAuthQueriesParams(req *http.Request, queries TransactionStatusQueries, v interface{}) error {

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token.Data.AccessToken))

	req.Header.Set("X-API-Key", c.TokenKey)

	return c.SendWithQueriesParams(req, queries, v)
}

func (c *Client) SendWithQueriesParams(req *http.Request, queries TransactionStatusQueries, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	qv := req.URL.Query()

	qv.Add("bankName", queries.BankName)
	qv.Add("pgReferenceId", queries.PgReferenceID)

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode == 500 {
		return errors.New("internal Server Error")

	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		data, err = io.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {

			return fmt.Errorf("error , status code: %d, %s", resp.StatusCode, string(data))

		}

		return fmt.Errorf("error reading body error , status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
