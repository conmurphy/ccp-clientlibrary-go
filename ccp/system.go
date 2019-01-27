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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LivenessHealth struct {
	CXVersion      *string `json:"CXVersion,omitempty"`
	TimeOnMgmtHost *string `json:"TimeOnMgmtHost,omitempty"`
}

type Health struct {
	TotalSystemHealth *string          `json:"TotalSystemHealth,omitempty"`
	CurrentNodes      *int64           `json:"CurrentNodes,omitempty"`
	ExpectedNodes     *int64           `json:"ExpectedNodes,omitempty"`
	NodesStatus       *[]NodeStatus    `json:"NodesStatus,omitempty"`
	PodStatusList     *[]PodStatusList `json:"PodStatusList,omitempty"`
}

type NodeStatus struct {
	NodeName           *string `json:"NodeName,omitempty"`
	NodeCondition      *string `json:"NodeCondition,omitempty"`
	NodeStatus         *string `json:"NodeStatus,omitempty"`
	LastTransitionTime *string `json:"LastTransitionTime,omitempty"`
}

type PodStatusList struct {
	PodName            *string `json:"PodName,omitempty"`
	PodCondition       *string `json:"PodCondition,omitempty"`
	PodStatus          *string `json:"PodStatus,omitempty"`
	LastTransitionTime *string `json:"LastTransitionTime,omitempty"`
}

func (s *Client) Login(client *Client) error {

	url := fmt.Sprintf(s.BaseURL + "/2/system/login?username=" + client.Username + "&password=" + client.Password)

	j, err := json.Marshal(client)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	_, err = s.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}

func (s *Client) GetLivenessHealth() (*LivenessHealth, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/system/livenessHealth")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data LivenessHealth

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	health := &data

	return health, nil
}

func (s *Client) GetHealth() (*Health, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/system/health")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Health

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	health := &data

	return health, nil
}
