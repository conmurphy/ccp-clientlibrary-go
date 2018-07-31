package ccp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Role struct {
	Role *string `json:"role,omitempty"`
}

func (s *Client) GetRole() (*Role, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/rbac")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Role

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	role := &data

	return role, nil
}
