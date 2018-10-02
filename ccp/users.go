package ccp

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (s *Client) GetUser(username string) (*User, error) {

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

	users := data

	for _, user := range users {

		if username == *user.Username {
			return &user, nil
		}
	}

	return nil, errors.New("USER NOT FOUND")
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

func (s *Client) PatchUser(user *User) (*User, error) {

	var data User

	if nonzero(user.Username) {
		return nil, errors.New("User.Username is missing")
	}

	username := *user.Username

	url := fmt.Sprintf(s.BaseURL + "/2/localusers/" + username)

	j, err := json.Marshal(user)

	if err != nil {

		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(j))
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

func (s *Client) DeleteUser(username string) error {

	if username == "" {
		return errors.New("Username of account to delete is required")
	}

	url := fmt.Sprintf(s.BaseURL + "/2/localusers/" + username)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
