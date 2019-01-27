/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

               https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.*/

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
