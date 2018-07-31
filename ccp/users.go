package ccp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/validator.v2"
)

//UserAPIResponse
type User struct {
	Token     *string `json:"Token,omitempty"`
	Username  *string `json:"UserName,omitempty" validate:"nonzero"`
	Disable   *bool   `json:"Disable,omitempty"`
	Role      *string `json:"Role,omitempty" validate:"nonzero"`
	FirstName *string `json:"FirstName,omitempty" `
	LastName  *string `json:"LastName,omitempty"`
	Password  *string `json:"Password,omitempty"`
}

func (s *Client) GetUsers() ([]User, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/localusers")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []User

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) AddUser(user *User) (*User, error) {

	var data User

	if errs := validator.Validate(user); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/2/localusers")

	j, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	user = &data

	return user, nil
}
