package ccp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LDAPSetup struct {
	Server                 *string `json:"Server,omitempty"`
	Port                   *int64  `json:"Port,omitempty" `
	BaseDN                 *string `json:"BaseDN,omitempty"`
	ServiceAccountDN       *string `json:"ServiceAccountDN,omitempty"`
	ServiceAccountPassword *string `json:"ServiceAccountPassword,omitempty"`
	StartTLS               *bool   `json:"StartTLS,omitempty"`
	InsecureSkipVerify     *bool   `json:"InsecureSkipVerify,omitempty" `
}

func (s *Client) GetLDAPSetup() (*LDAPSetup, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/ldap/setup")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *LDAPSetup

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
