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

type ACIProfile struct {
	UUID                     *string                    `json:"uuid,omitempty"`
	Name                     *string                    `json:"name,omitempty" `
	APICHosts                *string                    `json:"apic_hosts,omitempty"`
	APICUsername             *int64                     `json:"apic_username,omitempty"`
	APICPassword             *int64                     `json:"apic_password,omitempty"`
	ACIVMMDomainName         *string                    `json:"aci_vmm_domain_namestate,omitempty"`
	ACIInfraVLANID           *string                    `json:"aci_infra_vlan_id,omitempty" `
	VRFName                  *string                    `json:"vrf_name,omitempty"`
	L3OutsidePolicyName      *string                    `json:"l3_outside_policy_name,omitempty"`
	L3OutsideNetworkName     *string                    `json:"l3_outside_network_name,omitempty"`
	AAEPName                 *string                    `json:"aaep_name,omitempty"`
	Nameservers              *[]string                  `json:"nameservers,omitempty"`
	ACIAllocator             *ACIProfileAllocatorConfig `json:"aci_allocator,omitempty"`
	ControlPlaneContractName *string                    `json:"control_plane_contract_name,omitempty"`
}

type ACIProfileAllocatorConfig struct {
	NodeVLANStart      *int64  `json:"node_vlan_start,omitempty"`
	NodeVLANEnd        *int64  `json:"node_vlan_end,omitempty"`
	MulticastRange     *string `json:"multicast_range,omitempty"`
	ServiceSubnetStart *string `json:"service_subnet_start,omitempty"`
	PodSubnetStart     *string `json:"pod_subnet_start,omitempty"`
}

func (s *Client) GetACIProfiles() ([]ACIProfile, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/aci_profiles")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []ACIProfile

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
