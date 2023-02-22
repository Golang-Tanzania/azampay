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

func (api *APICONTEXT) GenerateSessionID(mode string) string {
	var authURL string
	if strings.ToLower(mode) == "production" {
		api.BaseURL = ProductionBaseURL
		authURL = ProductionAuthURL
	} else {
		api.BaseURL = SandboxBaseURL
		authURL = SandboxAuthURL
	}

	parameters := fmt.Sprintf(`{"appName":"%v", "clientId": "%v", "clientSecret": "%v"}`, api.AppName, api.ClientID, api.ClientSecret)

	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer([]byte(parameters)))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err.Error()
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err.Error()
	}

	type Result struct {
		Data       map[string]string
		Message    string
		Success    bool
		StatusCode int
	}

	var result Result

	json.Unmarshal([]byte(body), &result)

	if result.Data["accessToken"] != "" {
		api.Bearer = result.Data["accessToken"]
		fmt.Println(api.Bearer)
		return api.Bearer
	} else {
		return string(body)
	}
}

func (api *APICONTEXT) LoadKeys(file string) *APICONTEXT {

	keys, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	var Keys kEYS

	err = json.Unmarshal(keys, &Keys)

	if err != nil {
		fmt.Println(err)
	}

	api.AppName = Keys.AppName
	api.ClientID = Keys.ClientId
	api.ClientSecret = Keys.ClientSecret
	api.Token = Keys.Token

	return api
}
