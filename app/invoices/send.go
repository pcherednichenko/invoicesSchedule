package invoices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	errBadCodeTpl = "bad code in response: %d"
)

type messageRequest struct {
	Email string `json:"email"`
	Text  string `json:"text"`
}

type response struct {
	Paid    bool   `json:"paid"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// sendMessage with customer email and test via post request
func sendMessage(url string, email, text string) (bool, error) {
	reqBody := messageRequest{email, text}
	marshallBody, err := json.Marshal(reqBody)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshallBody))
	if err != nil {
		return false, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return false, fmt.Errorf(errBadCodeTpl, resp.StatusCode)
	}
	var respBody response
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return false, err
	}

	if respBody.Code != 0 && respBody.Code != http.StatusOK {
		return false, fmt.Errorf(errBadCodeTpl, respBody.Code)
	}
	return respBody.Paid, nil
}
