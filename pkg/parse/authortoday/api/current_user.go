package api

import "encoding/json"

func (a *HttpApi) FetchCurrentUser(userToken string) (*CurrentUser, error) {
	body, err := makeRequest("v1/account/current-user", userToken)
	if err != nil {
		return nil, err
	}
	currentUser := &CurrentUser{}
	err = json.Unmarshal(body, currentUser)
	if err != nil {
		return nil, err
	}

	return currentUser, nil
}

type CurrentUser struct {
	Id int `json:"id"`
}
