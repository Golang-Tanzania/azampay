package GoAzam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// GenerateSessionID() generates a token that will enable
// access to the endpoints. It accepts a string which will
// be the mode of the app, either Sandbox or Production.
// The default is Sandbox. It will return an error if token
// generation was unsuccessful
func (api *APICONTEXT) GenerateSession(mode string) error {
	var authURL string
	if strings.ToLower(mode) == "production" {
		api.BaseURL = ProductionBaseURL
		authURL = ProductionAuthURL
	} else {
		api.BaseURL = SandboxBaseURL
		authURL = SandboxAuthURL
	}

	parameters := fmt.Sprintf(`{"appName":"%v", "clientId": "%v", "clientSecret": "%v"}`, api.appName, api.clientID, api.clientSecret)

	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer([]byte(parameters)))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	type Result struct {
		Data       map[string]string
		Message    string
		Success    bool
		StatusCode int
	}

	var result Result

	if resp.StatusCode == 200 {
		decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&result)

		if decodeErr != nil {
			if decodeErr == io.EOF {
				return fmt.Errorf("(Token Generation) Error: Server returned an empty body")
			}
			return decodeErr
		}

		api.Bearer = result.Data["accessToken"]
		api.Expiry = result.Data["expire"]
		return nil

	} else if resp.StatusCode == 400 {

		var badRequest *BadRequestError

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&badRequest); err != nil {
			return fmt.Errorf("(Token Generation) Error: decoding bad request error: %w", err)
		}

		return fmt.Errorf(badRequest.Error())
	} else if resp.StatusCode == 423 {
		var invalidDetail *InvalidDetail

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&invalidDetail); err != nil {
			return fmt.Errorf("(Token Generation) Error: decoding invalid detail error: %w", err)
		}

		return fmt.Errorf(invalidDetail.Error())

	} else if resp.StatusCode == 500 {

		return fmt.Errorf("(Token Generation) Internal Server Error: status code 500")

	} else {

		return fmt.Errorf("(Token Generation) Error: status code %d", resp.StatusCode)

	}

}

// A function to read keys from a config.json file.
// It will return an error if any.
func (api *APICONTEXT) LoadKeys(file string) error {

	configKeys, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	var readKeys keys

	err = json.Unmarshal(configKeys, &readKeys)

	if err != nil {
		return err
	}

	api.appName = readKeys.AppName
	api.clientID = readKeys.ClientId
	api.clientSecret = readKeys.ClientSecret
	api.token = readKeys.Token

	return nil
}
