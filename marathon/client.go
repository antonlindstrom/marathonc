package marathon

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Run Get-request against an url and return body
func GetBody(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send request to endpoint
	return SendRequest(req)
}

// Post request to an endpoint
func PostBody(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	// Set JSON header
	req.Header.Add("Content-Type", "application/json")

	// Send request to endpoint
	return SendRequest(req)
}

// Send request
func SendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
