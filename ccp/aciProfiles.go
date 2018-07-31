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
