package ccp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProviderClientConfig struct {
	UUID   *string `json:"uuid,omitempty"`
	Name   *string `json:"name,omitempty" `
	Type   *int64  `json:"type,omitempty"`
	Config *Config `json:"config,omitempty"`
}

type Config struct {
	IP       *string `json:"ip,omitempty"`
	Port     *int64  `json:"po	rt,omitempty" `
	Username *string `json:"username,omitempty"`
}

type Vsphere struct {
	Datacenters *[]string `json:"Datacenters,omitempty"`
	Clusters    *[]string `json:"Clusters,omitempty"`
	VMs         *[]string `json:"VMs,omitempty"`
	Networks    *[]string `json:"Networks,omitempty"`
	Datastores  *[]string `json:"Datastores,omitempty"`
	Pools       *[]string `json:"Pools,omitempty"`
}

func (s *Client) GetProviderClientConfigs() ([]ProviderClientConfig, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []ProviderClientConfig

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfig(clientUUID string) (*ProviderClientConfig, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *ProviderClientConfig

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigClusters(clientUUID string) ([]Cluster, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/clusters")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []Cluster

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenter(clientUUID string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenterClusters(clientUUID string, datacenter string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter/" + datacenter + "/cluster")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenterVMs(clientUUID string, datacenter string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter/" + datacenter + "/vm")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenterNetworks(clientUUID string, datacenter string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter/" + datacenter + "/network")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenterDatastores(clientUUID string, datacenter string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter/" + datacenter + "/datastore")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetProviderClientConfigVsphereDatacenterClusterPools(clientUUID string, datacenter string, cluster string) (*Vsphere, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/providerclientconfigs/" + clientUUID + "/vsphere/datacenter/" + datacenter + "/cluster/" + cluster + "/pool")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Vsphere

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
