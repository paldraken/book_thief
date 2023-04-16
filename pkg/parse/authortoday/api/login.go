package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type loginReq struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

type loginRes struct {
	Token   string    `json:"token"`
	Issued  time.Time `json:"issued"`
	Expires time.Time `json:"expires"`
}

func (a *HttpApi) ObtainingAccessToken(login, password string) (string, error) {

	if login == "" || password == "" {
		return DEFAULT_TOKEN, nil
	}

	client := &http.Client{}

	loginData, err := json.Marshal(loginReq{Login: login, Password: password})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", BASE_DOMAIN+"/v1/account/login-by-password", strings.NewReader(string(loginData)))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+DEFAULT_TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("login failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	lres := &loginRes{}
	json.Unmarshal(body, lres)

	if lres.Token == "" {
		return "", errors.New("empty token")
	}

	return lres.Token, nil
}
