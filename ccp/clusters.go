package ccp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/validator.v2"
)

//ClusterAPIResponse
type Cluster struct {
	UUID                     *string        `json:"uuid,omitempty"`
	ProviderClientConfigUUID *string        `json:"provider_client_config_uuid,omitempty" validate:"nonzero"`
	ACIProfileUUID           *string        `json:"aci_profile_uuid,omitempty"`
	Name                     *string        `json:"name,omitempty"  validate:"nonzero"`
	Description              *string        `json:"description,omitempty"`
	ResourcePool             *string        `json:"resource_pool,omitempty"  validate:"nonzero"`
	Networks                 *[]string      `json:"networks,omitempty"  validate:"nonzero"`
	Workers                  *int64         `json:"workers,omitempty"  validate:"nonzero"`
	VCPUs                    *int64         `json:"vcpus,omitempty"  validate:"nonzero"`
	Memory                   *int64         `json:"memory,omitempty"  validate:"nonzero"`
	Type                     *int64         `json:"type,omitempty"   validate:"nonzero"`
	Masters                  *int64         `json:"masters,omitempty"  validate:"nonzero"`
	Datacenter               *string        `json:"datacenter,omitempty"  validate:"nonzero"`
	Cluster                  *string        `json:"cluster,omitempty" validate:"nonzero"`
	Datastore                *string        `json:"datastore,omitempty"  validate:"nonzero"`
	State                    *string        `json:"state,omitempty"`
	Template                 *string        `json:"template,omitempty"  validate:"nonzero" `
	SSHUser                  *string        `json:"ssh_user,omitempty"  validate:"nonzero"`
	SSHPassword              *string        `json:"ssh_password,omitempty"`
	SSHKey                   *string        `json:"ssh_key,omitempty"   validate:"nonzero"`
	Labels                   *[]Label       `json:"labels,omitempty"`
	Nodes                    *[]Node        `json:"nodes,omitempty"`
	DeployerType             *string        `json:"deployer_type,omitempty"  validate:"nonzero"`
	Deployer                 *Deployer      `json:"deployer,omitempty" validate:"nonzero"`
	KubernetesVersion        *string        `json:"kubernetes_version,omitempty" validate:"nonzero"`
	ClusterEnvURL            *string        `json:"cluster_env_url,omitempty"`
	ClusterDashboardURL      *string        `json:"cluster_dashboard_url,omitempty"`
	NetworkPlugin            *NetworkPlugin `json:"network_plugin,omitempty" validate:"nonzero"`
	CCPPrivateSSHKey         *string        `json:"ccp_private_ssh_key,omitempty"`
	CCPPublicSSHKey          *string        `json:"ccp_public_ssh_key,omitempty"`
	NTPPools                 *[]string      `json:"ntp_pools,omitempty"`
	NTPServers               *[]string      `json:"ntp_servers,omitempty"`
	IsControlCluster         *bool          `json:"is_control_cluster,omitempty"`
	IsAdopt                  *bool          `json:"is_adopt,omitempty"`
	RegistriesSelfSigned     *[]string      `json:"registries_self_signed,omitempty"`
	RegistriesInsecure       *[]string      `json:"registries_insecure,omitempty"`
	RegistriesRootCA         *[]string      `json:"registries_root_ca,omitempty"`
	IngressVIPPoolID         *string        `json:"ingress_vip_pool_id,omitempty"`
	IngressVIPAddrID         *string        `json:"ingress_vip_addr_id,omitempty"`
	IngressVIPs              *[]string      `json:"ingress_vips,omitempty"`
	KeepalivedVRID           *int64         `json:"keepalived_vrid,omitempty"`
	HelmCharts               *[]HelmChart   `json:"helm_charts,omitempty"`
	MasterVIPAddrID          *string        `json:"master_vip_addr_id,omitempty"`
	MasterVIP                *string        `json:"master_vip,omitempty"`
	MasterMACAddresses       *[]string      `json:"master_mac_addresses,omitempty"`
	ClusterHealthStatus      *string        `json:"cluster_health_status,omitempty"`
	AuthList                 *[]string      `json:"AuthList,omitempty"`
}

type Label struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Node struct {
	UUID              *string   `json:"uuid,omitempty"`
	Name              *string   `json:"name,omitempty"`
	PublicIP          *string   `json:"public_ip,omitempty"`
	PrivateIP         *string   `json:"private_ip,omitempty"`
	IsMaster          *bool     `json:"is_master,omitempty"`
	State             *string   `json:"state,omitempty"`
	CloudInitData     *string   `json:"cloud_init_data,omitempty"`
	KubernetesVersion *string   `json:"kubernetes_version,omitempty"`
	ErrorLog          *string   `json:"error_log,omitempty"`
	Template          *string   `json:"template,omitempty"`
	MacAddresses      *[]string `json:"mac_addresses,omitempty"`
}

type Deployer struct {
	ProxyCMD     *string   `json:"proxy_cmd,omitempty"`
	ProviderType *string   `json:"provider_type,omitempty" validate:"nonzero"`
	Provider     *Provider `json:"provider,omitempty" validate:"nonzero"`
	IP           *string   `json:"ip,omitempty"`
	Port         *int64    `json:"port,omitempty"`
	Username     *string   `json:"username,omitempty"`
	Password     *string   `json:"password,omitempty"`
}

type NetworkPlugin struct {
	Name    *string `json:"name,omitempty"`
	Status  *string `json:"status,omitempty"`
	Details *string `json:"details,omitempty"`
}

type HelmChart struct {
	HelmChartUUID *string `json:"helmchart_uuid,omitempty"`
	ClusterUUID   *string `json:"cluster_UUID,omitempty"`
	ChartURL      *string `json:"chart_url,omitempty"`
	Name          *string `json:"name,omitempty"`
	Options       *string `json:"options,omitempty"`
}

type Provider struct {
	VsphereDataCenter         *string              `json:"vsphere_datacenter,omitempty"`
	VsphereDatastore          *string              `json:"vsphere_datastore,omitempty"`
	VsphereSCSIControllerType *string              `json:"vsphere_scsi_controller_type,omitempty"`
	VsphereWorkingDir         *string              `json:"vsphere_working_dir,omitempty"`
	VsphereClientConfigUUID   *string              `json:"vsphere_client_config_uuid,omitempty" validate:"nonzero"`
	ClientConfig              *VsphereClientConfig `json:"client_config,omitempty"`
}

type VsphereClientConfig struct {
	IP       *string `json:"ip,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (s *Client) GetClusters() ([]Cluster, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters")

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

func (s *Client) GetCluster(clusterName string) (*Cluster, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Cluster

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetClusterHealth(clusterUUID string) (*Cluster, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID + "/health")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Cluster

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetClusterAuthz(clusterUUID string) (*Cluster, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID + "/authz")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *Cluster

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetClusterDashboard(clusterUUID string) (*string, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID + "/dashboard")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	data := string(bytes)
	return &data, nil
}

func (s *Client) GetClusterEnv(clusterUUID string) (*string, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID + "/env")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	data := string(bytes)
	return &data, nil
}

func (s *Client) GetClusterHelmCharts(clusterUUID string) (*HelmChart, error) {

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID + "/helmcharts")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data *HelmChart

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) AddCluster(cluster *Cluster) (*Cluster, error) {

	var data Cluster

	if errs := validator.Validate(cluster); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/2/clusters")

	j, err := json.Marshal(cluster)

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

	cluster = &data

	return cluster, nil
}