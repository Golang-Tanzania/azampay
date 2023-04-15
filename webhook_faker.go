package azampay

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendFakeRequest This imitates mwanana server on how it sends mwapp updates
func SendFakeRequest(endpoint string, params Update) (*http.Response, error) {
	mar, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	read := bytes.NewReader(mar)

	req, err := http.NewRequest("POST", endpoint, read)
	if err != nil {
		log.Printf("Error in posting fake request: %s", err.Error())
		return nil, err
	}

	if err != nil {
		log.Printf("Error in ComputeHMAC: %s", err.Error())
		return nil, err
	}

	//Todo set Auth headers
	//req.Header.Set("", "")
	//req.Header.Set("", "")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	return client.Do(req)
}
