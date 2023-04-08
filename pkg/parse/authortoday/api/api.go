package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	DEFAULT_TOKEN = "guest"
	BASE_DOMAIN   = "https://api.author.today/"
)

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func makeRequest(path string, userToken string) ([]byte, error) {
	reqUrl := BASE_DOMAIN + path
	client := &http.Client{}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+userToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, handleApiError(body)
	}

	return body, nil
}

func handleApiError(body []byte) error {
	cErr := &apiError{}
	err := json.Unmarshal(body, cErr)
	if err != nil {
		return errors.New("fetch chapter error")
	}
	return fmt.Errorf("fetch cahpter error. Code %s. %s ", cErr.Code, cErr.Message)
}
