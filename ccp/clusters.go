package ccp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	validator "gopkg.in/validator.v2"
)

//ClusterAPIResponse
type Cluster struct {
	UUID                      *string         `json:"uuid,omitempty"`
	ProviderClientConfigUUID  *string         `json:"provider_client_config_uuid,omitempty" validate:"nonzero"`
	ACIProfileUUID            *string         `json:"aci_profile_uuid,omitempty"`
	Name                      *string         `json:"name,omitempty"  validate:"nonzero"`
	Description               *string         `json:"description,omitempty"`
	Networks                  *[]string       `json:"networks,omitempty"  validate:"nonzero"`
	Datacenter                *string         `json:"datacenter,omitempty"  validate:"nonzero"`
	Datastore                 *string         `json:"datastore,omitempty"  validate:"nonzero"`
	Cluster                   *string         `json:"cluster,omitempty" validate:"nonzero"`
	ResourcePool              *string         `json:"resource_pool,omitempty"  validate:"nonzero"`
	Workers                   *int64          `json:"workers,omitempty"  validate:"nonzero"`
	VCPUs                     *int64          `json:"vcpus,omitempty"  "`
	Memory                    *int64          `json:"memory,omitempty"  `
	Type                      *int64          `json:"type,omitempty"  `
	Masters                   *int64          `json:"masters,omitempty"  validate:"nonzero"`
	State                     *string         `json:"state,omitempty"`
	Template                  *string         `json:"template,omitempty"   `
	SSHUser                   *string         `json:"ssh_user,omitempty"  validate:"nonzero"`
	SSHPassword               *string         `json:"ssh_password,omitempty"`
	SSHKey                    *string         `json:"ssh_key,omitempty"   validate:"nonzero"`
	Labels                    *[]Label        `json:"labels,omitempty"`
	Nodes                     *[]Node         `json:"nodes,omitempty"`
	Deployer                  *Deployer       `json:"deployer,omitempty" validate:"nonzero"`
	KubernetesVersion         *string         `json:"kubernetes_version,omitempty" validate:"nonzero"`
	ClusterEnvURL             *string         `json:"cluster_env_url,omitempty"`
	ClusterDashboardURL       *string         `json:"cluster_dashboard_url,omitempty"`
	NetworkPlugin             *NetworkPlugin  `json:"network_plugin,omitempty" validate:"nonzero"`
	CCPPrivateSSHKey          *string         `json:"ccp_private_ssh_key,omitempty"`
	CCPPublicSSHKey           *string         `json:"ccp_public_ssh_key,omitempty"`
	NTPPools                  *[]string       `json:"ntp_pools,omitempty"`
	NTPServers                *[]string       `json:"ntp_servers,omitempty"`
	IsControlCluster          *bool           `json:"is_control_cluster,omitempty"`
	IsAdopt                   *bool           `json:"is_adopt,omitempty"`
	RegistriesSelfSigned      *[]string       `json:"registries_self_signed,omitempty"`
	RegistriesInsecure        *[]string       `json:"registries_insecure,omitempty"`
	RegistriesRootCA          *[]string       `json:"registries_root_ca,omitempty"`
	IngressVIPPoolID          *string         `json:"ingress_vip_pool_id,omitempty"`
	IngressVIPAddrID          *string         `json:"ingress_vip_addr_id,omitempty"`
	IngressVIPs               *[]string       `json:"ingress_vips,omitempty"`
	KeepalivedVRID            *int64          `json:"keepalived_vrid,omitempty"`
	HelmCharts                *[]HelmChart    `json:"helm_charts,omitempty"`
	MasterVIPAddrID           *string         `json:"master_vip_addr_id,omitempty"`
	MasterVIP                 *string         `json:"master_vip,omitempty"`
	MasterMACAddresses        *[]string       `json:"master_mac_addresses,omitempty"`
	ClusterHealthStatus       *string         `json:"cluster_health_status,omitempty"`
	AuthList                  *[]string       `json:"auth_list,omitempty"`
	IsHarborEnabled           *bool           `json:"is_harbor_enabled,omitempty" `
	HarborAdminServerPassword *string         `json:"harbor_admin_server_password,omitempty"`
	HarborRegistrySize        *string         `json:"harbor_registry_size,omitempty"`
	LoadBalancerIPNum         *int64          `json:"load_balancer_ip_num,omitempty"  validate:"nonzero" `
	IsIstioEnabled            *bool           `json:"is_istio_enabled,omitempty"   `
	WorkerNodePool            *WorkerNodePool `json:"worker_node_pool,omitempty"  validate:"nonzero" `
	MasterNodePool            *MasterNodePool `json:"master_node_pool,omitempty"  validate:"nonzero" `
	Infra                     *Infra          `json:"infra,omitempty"  validate:"nonzero" `
}

type Infra struct {
	Datacenter   *string   `json:"datacenter,omitempty"  validate:"nonzero"`
	Datastore    *string   `json:"datastore,omitempty"  validate:"nonzero"`
	Cluster      *string   `json:"cluster,omitempty" validate:"nonzero"`
	Networks     *[]string `json:"networks,omitempty"  validate:"nonzero"`
	ResourcePool *string   `json:"resource_pool,omitempty"  validate:"nonzero"`
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

type WorkerNodePool struct {
	VCPUs    *int64  `json:"vcpus,omitempty" validate:"nonzero"`
	Memory   *int64  `json:"memory,omitempty" validate:"nonzero"`
	Template *string `json:"template,omitempty" validate:"nonzero"`
}

type MasterNodePool struct {
	VCPUs    *int64  `json:"vcpus,omitempty" validate:"nonzero"`
	Memory   *int64  `json:"memory,omitempty" validate:"nonzero"`
	Template *string `json:"template,omitempty" validate:"nonzero"`
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

func (s *Client) AddClusterBasic(cluster *Cluster) (*Cluster, error) {

	/*

		This function was added in order to provide users a better experience with adding clusters. The list of required
		fields has been shortend with all defaults and computed values such as UUIDs to be automatically configured on behalf of the user.

		The following fields and values will be configured. The remainder to be specified by the user

		ProviderClientConfigUUID
		KubernetesVersion - default will be set to 1.10.1
		Type - default will be set to 1
		Deployer
			ProviderType will be set to "vsphere"
			Provider
				VsphereDataCenter - already specified as part of Cluster struct so will use this same value
				VsphereClientConfigUUID
				VsphereDatastore - already specified as part of Cluster struct so will use this same value
				VsphereWorkingDir - default will be set to /VsphereDataCenter/vm
		NetworkPlugin
			Name - default will be set to contiv-vpp
			Status - default will be set to ""
			Details - default will be set to "{\"pod_cidr\":\"192.168.0.0/16\"}"
		WorkerNodePool
			VCPUs - default will be set to 2
			Memory - default will be set to 16384
		MasterNodePool
			VCPUs - default will be set to 2
			Memory - default will be set to 8192

	*/

	var data Cluster

	// The following will configured the defaults for the cluster as specified above as well as check that the minimum
	// fields are provided

	if nonzero(cluster.Name) {
		return nil, errors.New("Cluster.Name is missing")
	}
	if nonzero(cluster.Datacenter) {
		return nil, errors.New("Cluster.Datacenter is missing")
	}
	if nonzero(cluster.Cluster) {
		return nil, errors.New("Cluster.Cluster is missing")
	}
	if nonzero(cluster.ResourcePool) {
		return nil, errors.New("Cluster.ResourcePool is missing")
	}
	if nonzero(cluster.SSHUser) {
		return nil, errors.New("Cluster.SSHUser is missing")
	}
	if nonzero(cluster.SSHKey) {
		return nil, errors.New("Cluster.SSHKey is missing")
	}
	if nonzero(cluster.Workers) {
		return nil, errors.New("Cluster.Workers is missing")
	}
	if nonzero(cluster.Masters) {
		return nil, errors.New("Cluster.Masters is missing")
	}
	if nonzero(cluster.IsHarborEnabled) {
		return nil, errors.New("Cluster.IsHarborEnabled is missing")
	}
	if nonzero(cluster.IsIstioEnabled) {
		return nil, errors.New("Cluster.IsIstioEnabled is missing")
	}
	if nonzero(cluster.Template) {
		return nil, errors.New("Cluster.Template is missing")
	}

	// Retrieve the provider client config UUID rather than have the user need to provide this themselves.
	// This is also built for a single provider client config and as of CCP 1.5 this wll be Vsphere
	providerClientConfigs, err := s.GetProviderClientConfigs()

	if err != nil {
		return nil, err
	}

	networkPlugin := NetworkPlugin{
		Name:    String("contiv-vpp"),
		Status:  String(""),
		Details: String("{\"pod_cidr\":\"192.168.0.0/16\"}"),
	}

	provider := Provider{
		VsphereDataCenter:       String(*cluster.Datacenter),
		VsphereDatastore:        String(*cluster.Datastore),
		VsphereClientConfigUUID: String(*providerClientConfigs[0].UUID),
		VsphereWorkingDir:       String("/" + *cluster.Datacenter + "/vm"),
	}

	deployer := Deployer{
		ProviderType: String("vsphere"),
		Provider:     &provider,
	}

	workerNodePool := WorkerNodePool{
		VCPUs:    Int64(2),
		Memory:   Int64(16384),
		Template: String(*cluster.Template),
	}

	masterNodePool := MasterNodePool{
		VCPUs:    Int64(2),
		Memory:   Int64(16384),
		Template: String(*cluster.Template),
	}

	// Since it returns a list we will use the UUID from the first element
	cluster.ProviderClientConfigUUID = String(*providerClientConfigs[0].UUID)
	cluster.KubernetesVersion = String("1.10.1")
	cluster.Type = Int64(1)
	cluster.NetworkPlugin = &networkPlugin
	cluster.Deployer = &deployer
	cluster.WorkerNodePool = &workerNodePool
	cluster.MasterNodePool = &masterNodePool

	// Need to reset the cluster level template to nil otherwise we receive the following error
	// "Cluster level template cannot be provided when master_node_pool and worker_node_pool are provided"
	cluster.Template = nil

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

func (s *Client) PatchCluster(cluster *Cluster) (*Cluster, error) {

	var data Cluster

	if nonzero(cluster.UUID) {
		return nil, errors.New("Cluster.UUID is missing")
	}

	clusterUUID := *cluster.UUID

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + clusterUUID)

	j, err := json.Marshal(cluster)

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

	cluster = &data

	return cluster, nil
}

func (s *Client) DeleteCluster(uuid string) error {

	if uuid == "" {
		return errors.New("Cluster UUID to delete is required")
	}

	url := fmt.Sprintf(s.BaseURL + "/2/clusters/" + uuid)

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
