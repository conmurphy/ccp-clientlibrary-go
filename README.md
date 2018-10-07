# Container Platform Go Client Library

This is a Go Client Library used for accessing Cisco Container Platform (CCP). 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco Container Platform 1.5 with Go version 1.10

Table of Contents
=================

  * [CCP Go Client Library](#ccp-go-client-library)
      * [Quick Start](#quick-start)
      * [Quick Start - Creation from JSON file](#quick-start---creation-from-json-file)
      * [Helper Functions](#helper-functions)
         * [Without helper function](#without-helper-function)
         * [With helper function](#with-helper-function)
         * [Available Helper Functions](#available-helper-functions)
      * [Reference](#reference)
         * [System](#system)
         * [Users](#users)
         * [Clusters](#clusters)
         * [ProviderClientConfigs](#providerclientconfigs)
         * [ACIProfiles](#aciprofiles)
         * [LDAP](#ldap)
         * [RBAC](#rbac)


Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc)

## Quick Start

```golang
package main

import "github.com/ccp-clientlibrary-go/ccp”

/*
  Define new CCP client
*/

client := ccp.NewClient("admin", ”password", "https://my-ccp-address.com")

/*
  Retrieve login
*/

err := client.Login(client)

if err != nil {
  fmt.Println(err)
}

/*
  Print Users
*/

users, err := client.GetUsers()

if err != nil {
  fmt.Println(err)
} else {
  for _, user := range users {
    fmt.Printf("%+v\n", *user.Username)
  }
}
```

## Quick Start - Creation from JSON file

For some situations it may be easier to have the configuration represented as JSON rather than conifguring individually as per the  examples below (e.g. AddCluster). In this scenario you can either build the JSON file yourself or monitor the API POST call for the JSON data sent to CCP. This can be achieved using the browsers built in developer tools. See the following document for screenshots of how to find the POST call in the Chrome Developer Tools.

[Screenshots](https://github.com/conmurphy/ccp-clientlibrary-go/blob/master/README-DEVELOPER-TOOLS.md)


Example JSON File - newCluster.json
```json
{
  "name": "myContainerPlatformCluster",
  "kubernetes_version": "1.10.1",
  "ssh_key": "ssh-rsa aaabbbmysshkey me@localhost",
  "description": "My first CCP Cluster",
  "datacenter": "innovation-lab",
  "cluster": "hx-cluster",
  "resource_pool": "hx-cluster/Resources",
  "datastore": "CCP",
  "ssh_user": "ccp",
  "template": "ccp-tenant-image-1.10.1-1.1.0.ova",
  "masters": 1,
  "workers": 2,
  "vcpus": 2,
  "memory": 16384,
  "type": 1,
  "ingress_vip_pool_id": "12345abcd-abcd1234-1234543221",
    "network_plugin": {
      "name": "contiv-vpp",
      "status": "",
      "details": "{\"pod_cidr\":\"192.168.0.0/16\"}"
    },
  "provider_client_config_uuid": "1234abcd-abcd1234-abcdabcd",
  "networks": ["ccp-network/ccp-network-port-group"],
  "deployer": {
    "provider_type": "vsphere",
    "provider": {
      "vsphere_datacenter": "innovation-lab",
      "vsphere_datastore": "CCP",
      "vsphere_client_config_uuid": "1234abcd-abcd1234-abcdabcd",
      "vsphere_working_dir": "/innovation-lab/vm"
    }
  }
}
```

```golang
package main

import (
  "fmt"
  "github.com/ccp-clientlibrary-go/ccp"
)



/*
  Define new ccp client
*/

client := ccp.NewClient("admin", ”password", "https://my-ccp-address.com")

/*
  Retrieve login
*/

err := client.Login(client)

if err != nil {
  fmt.Println(err)
}

/*
  Create cluster
*/
	
clusterJSONFile, err := os.Open("newCluster.json")

if err != nil {
	fmt.Println(err)
}

bytes, _ := ioutil.ReadAll(clusterJSONFile)

var cluster *ccp.Cluster

json.Unmarshal(bytes, &cluster)

cluster, err = client.AddCluster(cluster)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cluster UUID: " + *cluster.UUID)
}

defer clusterJSONFile.Close()
```

## Helper Functions

As per the following link, using the Marshal function from the encoding/json library treats false booleans as if they were nil values, and thus it omits them from the JSON response. To make a distinction between a non-existent boolean and false boolean we need to use a ```*bool``` in the struct. 

```golang
type User struct {
	FirstName               *string `json:"firstName,omitempty"`
	LastName                *string `json:"lastName,omitempty"`
	Password                *string `json:"password,omitempty"` 
}
```
https://github.com/golang/go/issues/13284

Therefore in order to have a consistent experience all struct fields within this client library use pointers. This provides a way to differentiate between unset values, nil, and an intentional zero value, such as "", false, or 0. 

Helper functions have been created to simplify the creation of pointer types.

### Without helper function

```golang
firstName 	:= "client"
lastName 	:= "library"
password	:= "myPassword"

newUser := ccp.User {
	FirstName:   &firstName,
	LastName:    &lastName,
	Password:    &password,
}
```
### With helper function

```golang
newUser := ccp.User {
	FirstName:   ccp.String("client"),
	LastName:    ccp.String("library"),
	Password:    ccp.String("myPassword"),
}
```

Reference: https://willnorris.com/2014/05/go-rest-apis-and-pointers

### Available Helper Functions

* ccp.Bool()
* ccp.Int()
* ccp.Int64()
* ccp.String()
* ccp.Float32()
* ccp.Float64()

## Reference

- [System](#system)
- [Users](#users)
- [Clusters](#clusters)
- [ProviderClientConfigs](#providerclientconfigs)
- [ACIProfiles](#aciprofiles)
- [LDAP](#ldap)
- [RBAC](#rbac)

### System

- [Login](#login)
- [GetLivenessHealth](#getlivenesshealth)
- [GetHealth](#gethealth)

```go
type LivenessHealth struct {
	CXVersion      *string 
	TimeOnMgmtHost *string
}
```

```go
type Health struct {
	TotalSystemHealth *string          
	CurrentNodes      *int64           
	ExpectedNodes     *int64           
	NodesStatus       *[]NodeStatus    
	PodStatusList     *[]PodStatusList 
}
```

```go
type NodeStatus struct {
	NodeName           *string 
	NodeCondition      *string 
	NodeStatus         *string 
	LastTransitionTime *string 
}
```

```go
type PodStatusList struct {
	PodName            *string 
	PodCondition       *string
	PodStatus          *string
	LastTransitionTime *string 
}
```

#### Login

```go
func (s *Client) Login(client *Client) error
```

##### Example

```go
client := ccp.NewClient("admin", ”password", "https://my-ccp-address.com")

err := client.Login(client)

if err != nil {
	fmt.Println(err)
}
```

#### GetLivenessHealth

```go
func (s *Client) GetLivenessHealth() (*LivenessHealth, error)
```

##### Example

```go

```

#### GetHealth

```go
func (s *Client) GetHealth() (*Health, error)
```

##### Example
```go

```

### Users

[Field Explanations](#field-explanations)

- [GetUsers](#getusers)
- [GetUser](#getuser)
- [AddUser](#adduser)
- [PatchUser](#patchuser)
- [DeleteUser](#deleteuser)

```go
type User struct {
	Username  *string 
	Disable   *bool  
	Role      *string 
	FirstName *string
	LastName  *string
	Password  *string
}
```

#### Field Explanations

Field | Description 
------------ | -------------
Role | Role of the user - either Administrator or Devops
Disable | Whether or not the user account is enabled or disabled
	
	
#### GetUsers

```go
func (s *Client) GetUsers() ([]User, error)
```

##### Example
```go  
  users, err := client.GetUsers()
  
  if err != nil {
    fmt.Println(err)
  } else {
    for _, user := range users {
      fmt.Printf("%+v\n", *user.Username)
    }
  }
```

#### GetUser

```go
func (s *Client) GetUser(username string) (*User, error)
```

##### Example
```go  
user, err := client.GetUser("myUsername")
  
if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", *user.Username)
  fmt.Printf("%+v\n", *user.Role)
}
```

#### AddUser

```go
func (s *Client) AddUser(user *User) (*User, error) {
```

##### __Required Fields__
* Username
* Role

  
##### Example
```go
newUser := ccp.User{
  FirstName: ccp.String("ccp"),
  LastName:  ccp.String("sdk"),
  Username:  ccp.String("ccp_sdk"),
  Password:  ccp.String("password123"),
  Disable:   ccp.Bool(false),
  Role:      ccp.String("SysAdmin"),
}

user, err := client.AddUser(&newUser)

if err != nil {
  fmt.Println(err)
} else {
  username := *user.Username
  token := *user.Token
  fmt.Println("Username: " + username + ", Token: " + token)
}
```

#### PatchUser

```go
func (s *Client) PatchUser(user *User) (*User, error) 
```

##### __Required Fields__
* Username

##### __Available Fields to Patch__
* Firstname
* LastName
* Password
* Disable
* Role
	
  
##### Example
```go
newUser := ccp.User{
  Username:  ccp.String("ccp_sdk"),
  Role:      ccp.String("Devops"),
}

user, err := client.PatchUser(&newUser)

if err != nil {
  fmt.Println(err)
} else {
  username := *user.Username
  role := *user.Role
  fmt.Println("Username: " + username + ", Role: " + role)
}
```

#### DeleteUser

```go
func (s *Client) DeleteUser(username string) error 
```
  
##### Example
```go
err := client.DeleteUser("ccp_sdk")

if err != nil {
  fmt.Println(err)
}
```

### Clusters

[Field Explanations](#field-explanations)

- [GetClusters](#getclusters)
- [GetCluster](#getcluster)
- [GetClusterHealth](#getclusterhealth)
- [GetClusterAuthz](#getclusterauthz)
- [GetClusterDashboard](#getclusterdashboard)
- [GetClusterEnv](#getclusterenv)
- [GetClusterHelmCharts](#getclusterhelmcharts)
- [AddCluster](#addcluster)
- [PatchCluster](#patchcluster)
- [DeleteCluster](#deletecluster)

```go
type Cluster struct {
	UUID                       *string  
	ProviderClientConfigUUID   *string  
	ACIProfileUUID             *string 
	Name                       *string  
	Description                *string   
	Workers                    *int64    
	Masters                    *int64   
	ResourcePool               *string          
	Networks                   *[]string 
	Type                       *int64 
	Datacenter                 *string 
	Cluster                    *string        
	Datastore                  *string 
	State                      *string 
	Template                   *string 
	SSHUser                    *string 
	SSHPassword                *string 
	SSHKey                     *string 
	Labels                     *[]Label 
	Nodes                      *[]Node   
	Deployer                   *KubeADM              
	KubernetesVersion          *string               
	ClusterEnvURL              *string               
	ClusterDashboardURL        *string               
	NetworkPlugin              *NetworkPlugin
	CCPPrivateSSHKey           *string              
	CCPPublicSSHKey            *string              
	NTPPools                   *[]string       
	NTPServers                 *[]string      
	IsControlCluster           *bool             
	IsAdopt                    *bool              
	RegistriesSelfSigned       *[]string           
	RegistriesInsecure         *[]string            
	RegistriesRootCA           *[]string          
	IngressVIPPoolID           *string             
	IngressVIPAddrID           *string              
	IngressVIPs                *[]string             
	KeepalivedVRID             *int64              
	HelmCharts                 *[]HelmChart    
	MasterVIPAddrID            *string          
	MasterVIP                  *string        
	MasterMACAddresses         *[]string      
	ClusterHealthStatus        *string       
	AuthList                   *[]string 
	IsHarborEnabled            *bool           
	HarborAdminServerPassword  *string        
	HarborRegistrySize         *string        
	LoadBalancerIPNum          *int64          
	IsIstioEnabled             *bool          
	WorkerNodePool             *WorkerNodePool  
	MasterNodePool             *MasterNodePool  
	Infra                      *Infra 
}

type Infra struct {
	Datacenter   *string   
	Datastore    *string  
	Cluster      *string   
	Networks     *[]string
	ResourcePool *string   
}

type Label struct {
	Key                        *string  
	Value                      *string  
}

type Node struct {
	UUID                       *string   
	Name                       *string   
	PublicIP                   *string    
	PrivateIP     		   *string   
	IsMaster     		   *bool  
	State     	           *string   
	CloudInitData  		   *string    
	KubernetesVersion          *string   
	ErrorLog         	   *string   
	Template       	           *string   
	MacAddresses               *[]string  
}

type Deployer struct {
	ProxyCMD     *string    
	ProviderType *string   
	Provider     *Provider 
	IP           *string  
	Port         *int64   
	Username     *string    
	Password     *string    

type NetworkPlugin struct {
	Name   			   *string  
	Status 			   *string  
	Details			   *string  
}

type HelmChart struct {
	HelmChartUUID		   *string  
	ClusterUUID  		   *string  
	ChartURL     		   *string  
	Name         		   *string  
	Options     		   *string  
}	

type Provider struct {
	VsphereDataCenter          *string             
	VsphereDatastore           *string             
	VsphereSCSIControllerType  *string           
	VsphereWorkingDir          *string           
	VsphereClientConfigUUID    *string          
	ClientConfig               *VsphereClientConfig  
}

type VsphereClientConfig struct {
	IP       		   *string  
	Port     		   *int64  
	Username 		   *string  
	Password 		   *string  
}

type WorkerNodePool struct {
	VCPUs   		   *int64   
	Memory  		   *int64   
	Template		   *string  
}

type MasterNodePool struct {
	VCPUs    		   *int64   
	Memory   		   *int64   
	Template 		   *string  
}
```

#### Field Explanations

Type | Field | Description 
------------ | ------------ | -------------
Cluster	|	UUID	|	UUID of the  cluster 
Cluster	|	ProviderClientConfigUUID	|	UUID of the provider for the cluster (e.g. vsphere provider) which can be found using the ```GetProviderClientConfigs()``` function
Cluster	|	ACIProfileUUID	|	UUID of the ACI profile used with the cluster which can be found using the  ```GetACIProfiles()``` function
Cluster	|	Name	|	Name of the new cluster
Cluster	|	Description	|	Description for the new cluster
Cluster	|	Workers	|	Number of worker nodes. Must be greater than 0
Cluster	|	Masters	|	Number of master nodes. As of release 1.5 this value should be 1
Cluster	|	ResourcePool	|	The Vsphere resource pool in which the nodes will be running. If no resources have been created this is typically ```[cluster-name]/Resources```    
Cluster	|	Networks	|	Networks that the nodes will use, in the case of Vsphere these will be the names of the port groups that will attach to the K8s nodes. If using Hyperflex remember to include the ```k8-priv-iscsivm-network```      
Cluster	|	Type	|	As of CCP 1.5 this should be set to 1
Cluster	|	Datacenter	|	Vsphere datacenter in which the nodes will be deployed
Cluster	|	Cluster	|	Vsphere cluster on which the nodes will be deployed      
Cluster	|	Datastore	|	Vsphere datastore on which the nodes will be deployed      
Cluster	|	Template	|	The Vsphere template from which the nodes will be deployed. This should have been deployed at the initial installation e.g. ccp-tenant-image-1.10.1-ubuntu16-1.5.0   
Cluster	|	SSHUser	|	Username of a user to setup on each of the nodes as part of the cluster deployment. The nodes will then be accessible using this username and SSH key below. Use case includes troubleshooting
Cluster	|	SSHPassword	|	Password for the SSH user specified above
Cluster	|	SSHKey	|	Key for the SSH user specified above
Cluster	|	Labels	|	Labels configuration - See below
Cluster	|	Nodes	|	Node configuration - See below
Cluster	|	Deployer	|	Deployer configuration - See below
Cluster	|	Kubernetes Version	|	Version of Kubeternes to use
Cluster	|	ClusterEnvURL	|	
Cluster	|	ClusterDashboardURL	|	
Cluster	|	NetworkPlugin	|	
Cluster	|	CCPPrivateSSHKey	|	
Cluster	|	CCPPublicSSHKey	|	
Cluster	|	NTPPools	|	
Cluster	|	NTPServers	|	
Cluster	|	IsControlCluster	|	
Cluster	|	IsAdopt	|	
Cluster	|	RegistriesSelfSigned	|	
Cluster	|	RegistriesInsecure	|	
Cluster	|	RegistriesRootCA	|	
Cluster	|	IngressVIPPoolID	|	Required if using Load Balancer IP
Cluster	|	IngressVIPAddressID	|	
Cluster	|	IngressVIPs	|	
Cluster	|	KeepaliveVRID	|	
Cluster	|	HelmCharts	|	
Cluster	|	MasterVIPAddressID	|	
Cluster	|	MasterVIP	|	
Cluster	|	MasterMACAddresses	|	
Cluster	|	ClusterHealthStatus	|	
Cluster	|	AuthList	|	
Cluster	|	IsHarborEnabled	|	Whether or not Harbor is enabled- True or False
Cluster	|	HarborAdminServerPassword	|	
Cluster	|	HarborRegistrySize	|	
Cluster	|	LoadBalancerIPNum	|	Number of IP addresses to use from the VIP pool. If Istio is enabled this should be 3 or greater
Cluster	|	IsIstioEnabled	|	Whether or not Istio is enabled - True or False
Cluster	|	WorkerNodePool	|	Worker Node configuration - See below 
Cluster	|	MasterNodePool	|	Master Node configuration - See below 
Infra	|	Datacenter	|	Vsphere datacenter in which the nodes will be deployed
Infra	|	Datastore	|	Vsphere cluster on which the nodes will be deployed      
Infra	|	Cluster	|	Vsphere datastore on which the nodes will be deployed      
Infra	|	Networks	|	Networks that the nodes will use, in the case of Vsphere these will be the names of the port groups that will attach to the K8s nodes. If using Hyperflex remember to include the ```k8-priv-iscsivm-network```      
Infra	|	ResourcePool	|	The Vsphere resource pool in which the nodes will be running. If no resources have been created this is typically ```[cluster-name]/Resources```    
Label	|	Key	|	
Label	|	Value	|	
Node	|	UUID	|	
Node	|	Name	|	
Node	|	PublicIP	|	
Node	|	PrivateIP	|	
Node	|	IsMaster	|	
Node	|	State	|	
Node	|	CloudInitData	|	
Node	|	KubernetesVersion	|	Version of Kubeternes running
Node	|	ErrorLog	|	
Node	|	Template	|	
Node	|	MacAddresses	|	
Deployer	|	ProxyCMD	|	
Deployer	|	ProviderType	|	
Deployer	|	Provider	|	
Deployer	|	IP	|	
Deployer	|	Port	|	
Deployer	|	Username	|	
Deployer	|	Password	|	
NetworkPlugin	|	Name	|	
NetworkPlugin	|	Status	|	
NetworkPlugin	|	Details	|	
HelmChart	|	HelmChartUUID	|	
HelmChart	|	ClusterUUID	|	
HelmChart	|	ChartURL	|	
HelmChart	|	Name	|	
HelmChart	|	Options	|	
Provider	|	VsphereDataCenter	|	
Provider	|	VsphereDatastore	|	
Provider	|	VsphereSCSIControllerType	|	
Provider	|	VsphereWorkingDir	|	
Provider	|	VsphereClientConfigUUID	|	
Provider	|	ClientConfig	|	
VsphereClientConfig	|	IP	|	
VsphereClientConfig	|	Port	|	
VsphereClientConfig	|	Username	|	
VsphereClientConfig	|	Password	|	
WorkerNodePool	|	VCPUs	|	Amount of vCPUs each K8s worker node will use
WorkerNodePool	|	Memory	|	Amount of memory each K8s worker node will use
WorkerNodePool	|	Template	|	The Vsphere template from which the nodes will be deployed. This should have been deployed at the initial installation e.g. ccp-tenant-image-1.10.1-ubuntu16-1.5.0   
MasterNodePool	|	VCPUs	|	Amount of vCPUs each K8s master node will use
MasterNodePool	|	Memory	|	Amount of memory each K8s master node will use
MasterNodePool	|	Template	|	The Vsphere template from which the nodes will be deployed. This should have been deployed at the initial installation e.g. ccp-tenant-image-1.10.1-ubuntu16-1.5.0   	
	
#### GetClusters

```go
func (s *Client) GetClusters() ([]Cluster, error)
```

##### Example
```go  
  cluster, err := client.GetClusters()
  
  if err != nil {
    fmt.Println(err)
  } else {
    for _, cluster := range clusters {
      fmt.Printf("%+v\n", *cluster.Name)
    }
  }
```

#### GetCluster

```go
func (s *Client) GetCluster(clusterName string) (*Cluster, error)
```

##### Example
```go
  cluster, err := client.GetCluster("myCluster")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *cluster.UUID)
  }
```

#### GetClusterHealth

```go
func (s *Client) GetClusterHealth(clusterUUID string) (*Cluster, error) 
```

##### Example
```go

```

#### GetClusterAuthz

```go
func (s *Client) GetClusterAuthz(clusterUUID string) (*Cluster, error)
```

##### Example
```go
  clusterAuthz, err := client.GetClusterAuthz("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *clusterAuthz.AuthList)
  }
```

### GetClusterDashboard

```go
func (s *Client) GetClusterDashboard(clusterUUID string) (*string, error)
```

##### Example
```go
  clusterDashboardAddress, err := client.GetClusterDashboard("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *clusterDashboardAddress)
  }
```

### GetClusterEnv

```go
func (s *Client) GetClusterEnv(clusterUUID string) (*string, error) 
```

##### Example
```go
  clusterEnvironment, err := client.GetClusterEnv("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *clusterEnvironment)
  }
```

### GetClusterHelmCharts

```go
func (s *Client) GetClusterHelmCharts(clusterUUID string) (*HelmChart, error)
```

##### Example
```go
  clusterHelmCharts, err := client.GetClusterHelmCharts("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
    for _, clusterHelmChart := range clusterHelmCharts {
      fmt.Printf("%+v\n", *clusterHelmChart.Name)
    }
  }
```

#### AddCluster

```go
func (s *Client) AddCluster(cluster *Cluster) (*Cluster, error)
```

##### __Required Fields__
* ProviderClientConfigUUID
* Name
* KubernetesVersion
* ResourcePool
* Networks
* SSHKey
* Datacenter
* Cluster
* Datastore
* Workers
* SSHUser
* Type
* Masters
* Deployer
  * ProviderType
  * Provider 
    * VsphereDataCenter
    * VsphereClientConfigUUID
    * VsphereDatastore
    * VsphereWorkingDir
* NetworkPlugin
  * Name 
  * Status
  * Details
* IsHarborEnabled         
* LoadBalancerIPNum                
* IsIstioEnabled             
* WorkerNodePool    
  * VCPUs    
  * Memory  
  * Template 
* MasterNodePool           
  * VCPUs    
  * Memory  
  * Template 
  
##### Example
```go

workerNodePool := ccp.WorkerNodePool{
  VCPUs:    ccp.Int64(2),
  Memory:  ccp.Int64(16384),
  Template: ccp.String("ccp-tenant-image-1.10.1-1.4.0"),
}

masterNodePool := ccp.MasterNodePool{
  VCPUs:    ccp.Int64(2),
  Memory:  ccp.Int64(16384),
  Template: ccp.String("ccp-tenant-image-1.10.1-1.4.0"),
}
 
networkPlugin := ccp.NetworkPlugin{
  Name:    ccp.String("contiv-vpp"),
  Status:  ccp.String(""),
  Details: ccp.String("{\"pod_cidr\":\"192.168.0.0/16\"}"),
}
	
provider := ccp.Provider{
  VsphereDataCenter:       ccp.String("ccp-lab"),
  VsphereDatastore:        ccp.String("ccpDatastore"),
  VsphereClientConfigUUID: ccp.String("example-uuid-aaa-bbb-ccc"),
  VsphereWorkingDir:       ccp.String("/ccp-lab/vm"),
}

deployer := ccp.Deployer{
  ProviderType: ccp.String("vsphere"),
  Provider: &provider,
}

var networks []string

networks = append(networks, "ccp-network/ccp-network-portgroup")
	
newCluster := ccp.Cluster{
  ProviderClientConfigUUID: ccp.String("1234abcd-1234-0000-aaaa-abcdef12345"),
  Name:                     ccp.String("ccp-api-cluster"),
  KubernetesVersion:        ccp.String("1.10.1"),
  SSHKey:            	    ccp.String("ssh-rsa sshkey123abc me@locahost"),
  Datacenter:       	    ccp.String("ccp-lab"),
  Cluster:                  ccp.String("hx-cluster"),
  ResourcePool: 	    ccp.String("hx-cluster/Resources"),
  Networks:    		    &networks,
  Datastore:    	    ccp.String("ccpDatastore"),
  Template:     	    ccp.String("ccp-tenant-image-1.10.1-1.1.0.ova"),
  Masters:      	    ccp.Int64(1),
  Workers:      	    ccp.Int64(2),
  SSHUser:      	    ccp.String("ccpuser"),
  Type:         	    ccp.Int64(1),
  Deployer: 		    &deployer,
  NetworkPlugin:            &networkPlugin,
  IsHarborEnabled: 	    ccp.Bool(false),	    
  LoadBalanderIPNum: 	    ccp.Int64(1),                
  IsIstioEnabled: 	    ccp.Bool(false),
  WorkerNodePool:           &workerNodePool,
  MasterNodePool:           &masterNodePool,
}

cluster, err := client.AddCluster(&newCluster)

if err != nil {
  fmt.Println(err)
} else {
  fmt.Println("Cluster UUID: " + *cluster.UUID)
}
 
```
#### PatchCluster

```go
func (s *Client) PatchCluster(cluster *Cluster) (*Cluster, error) 
```

##### __Required Fields__
* UUID
* Workers 

##### __Available Fields To Patch__
* Workers
* LoadBalanderIPNum
  
##### Example
```go

newCluster := ccp.Cluster{
  UUID: ccp.String("aaaa-bbbb-cccc-dddd-eeee"),
  Workers: ccp.Int64(3),
  LoadBalanderIPNum: ccp.Int64(3),
}	
cluster, err := client.PatchCluster(&newCluster)

if err != nil {
  fmt.Println(err)
} else {
  fmt.Println("Cluster UUID: " + *cluster.UUID)
}
 
```

### DeleteCluster

```go
func (s *Client) DeleteCluster(uuid string) error 
```

##### Example
```go
err = client.DeleteCluster("aaaa-bbbb-cccc-dddd-eeee")

if err != nil {
  fmt.Println(err)
}
```

### ProviderClientConfigs

- [GetProviderClientConfigs](#getproviderclientconfigs)
- [GetProviderClientConfig](#getproviderclientconfig)
- [GetProviderClientConfigClusters](#getproviderclientconfigclusters)
- [GetProviderClientConfigVsphereDatacenter](#getproviderclientconfigvspheredatacenter)
- [GetProviderClientConfigVsphereDatacenterClusters](#getproviderclientconfigvspheredatacenterclusters)
- [GetProviderClientConfigVsphereDatacenterVMs](#getproviderclientconfigvspheredatacentervms)
- [GetProviderClientConfigVsphereDatacenterNetworks](#getproviderclientconfigvspheredatacenternetworks)
- [GetProviderClientConfigVsphereDatacenterDatastores](#getproviderclientconfigvspheredatacenterdatastores)
- [GetProviderClientConfigVsphereDatacenterClusterPools](#getproviderclientconfigvspheredatacenterclusterpools)

```go
type ProviderClientConfig struct {
	UUID   		*string  
	Name   		*string  
	Type   		*int64 
	Config 		*Config  
}

type Config struct {
	IP       	*string  
	Port     	*int64  
	Username 	*string  
}

type Vsphere struct {
	Datacenters 	*[]string  
	Clusters    	*[]string 
	VMs         	*[]string  
	Networks    	*[]string  
	Datastores  	*[]string 
	Pools       	*[]string  
}
```

### GetProviderClientConfigs

```go
func (s *Client) GetProviderClientConfigs() ([]ProviderClientConfig, error)
```

##### Example
```go
  providerClientConfigs, err := client.GetProviderClientConfigs()
  
  if err != nil {
    fmt.Println(err)
  } else {
    for _, providerClientConfig := range providerClientConfigs {
      fmt.Printf("%+v\n", *providerClientConfig.Name)
    }
  }
```

### GetProviderClientConfig

```go
func (s *Client) GetProviderClientConfig(clientUUID string) (*ProviderClientConfig, error)
```

##### Example
```go
  providerClientConfig, err := client.GetProviderClientConfig("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%+v\n", *providerClientConfig.Name)
  }
```

### GetProviderClientConfigClusters

```go
func (s *Client) GetProviderClientConfigClusters(clientUUID string) ([]Cluster, error)
```

##### Example
```go
  providerClientConfigClusters, err := client.GetProviderClientConfigClusters("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
     for _, providerClientConfigCluster := range providerClientConfigClusters {
      fmt.Printf("%+v\n", *providerClientConfigCluster.Name)
    }
  }
```

### GetProviderClientConfigVsphereDatacenter

```go
func (s *Client) GetProviderClientConfigVsphereDatacenter(clientUUID string) (*Vsphere, error) 
```

##### Example
```go
  providerClientConfigVsphereDatacenter, err := client.GetProviderClientConfigVsphereDatacenter("AAAA-BBBB-CCCC-UUID")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenter.Datacenters)
  }
```

### GetProviderClientConfigVsphereDatacenterClusters

```go
func (s *Client) GetProviderClientConfigVsphereDatacenterClusters(clientUUID string, datacenter string) (*Vsphere, error)
```

##### Example
```go
  providerClientConfigVsphereDatacenterClusters, err := client.GetProviderClientConfigVsphereDatacenterClusters("AAAA-BBBB-CCCC-UUID", "myDatacenter")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenterClusters.Clusters)
  }
```

### GetProviderClientConfigVsphereDatacenterVMs

```go
func (s *Client) GetProviderClientConfigVsphereDatacenterVMs(clientUUID string, datacenter string) (*Vsphere, error)
```

##### Example
```go
  providerClientConfigVsphereDatacenterVMs, err := client.GetProviderClientConfigVsphereDatacenterVMs("AAAA-BBBB-CCCC-UUID", "myDatacenter")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenterVMs.VMs)
  }
```

### GetProviderClientConfigVsphereDatacenterNetworks

```go
func (s *Client) GetProviderClientConfigVsphereDatacenterNetworks(clientUUID string, datacenter string) (*Vsphere, error)
```

##### Example
```go
  providerClientConfigVsphereDatacenterNetworks, err := client.GetProviderClientConfigVsphereDatacenterNetworks("AAAA-BBBB-CCCC-UUID", "myDatacenter")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenterNetworks.Networks)
  }
```

### GetProviderClientConfigVsphereDatacenterDatastores

```go
func (s *Client) GetProviderClientConfigVsphereDatacenterDatastores(clientUUID string, datacenter string) (*Vsphere, error)
```

##### Example
```go
  providerClientConfigVsphereDatacenterDatastores, err := client.GetProviderClientConfigVsphereDatacenterDatastores("AAAA-BBBB-CCCC-UUID", "myDatacenter")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenterDatastores.Datastores)
  }
```

### GetProviderClientConfigVsphereDatacenterClusterPools

```go
func (s *Client) GetProviderClientConfigVsphereDatacenterClusterPools(clientUUID string, datacenter string, cluster string) (*Vsphere, error) 
```

##### Example
```go
  providerClientConfigVsphereDatacenterPools, err := client.GetProviderClientConfigVsphereDatacenterClusterPools("AAAA-BBBB-CCCC-UUID", "myDatacenter", "myCluster")
  
  if err != nil {
    fmt.Println(err)
  } else {
      fmt.Printf("%+v\n", *providerClientConfigVsphereDatacenterPools.Pools)
  }
```

### ACIProfiles

- [GetACIProfiles](#getaciprofiles)


```go
type ACIProfile struct {
	UUID                   	   *string                
	Name                 	   *string               
	APICHosts              	   *string                
	APICUsername               *int64                
	APICPassword               *int64               
	ACIVMMDomainName           *string           
	ACIInfraVLANID             *string           
	VRFName                    *string      
	L3OutsidePolicyName        *string         
	L3OutsideNetworkName       *string         
	AAEPName                   *string              
	Nameservers                *[]string             
	ACIAllocator               *ACIProfileAllocatorConfig 
	ControlPlaneContractName   *string                     
}

type ACIProfileAllocatorConfig struct {
	NodeVLANStart     	   *int64   
	NodeVLANEnd       	   *int64  
	MulticastRange     	   *string  
	ServiceSubnetStart 	   *string 
	PodSubnetStart     	   *string  
}
```

### GetACIProfiles

```go
func (s *Client) GetACIProfiles() ([]ACIProfile, error) 
```

##### Example
```go
  aciProfiles, err := client.GetACIProfiles()
  
  if err != nil {
    fmt.Println(err)
  } else {
    for _, aciProfile := range aciProfiles {
      fmt.Printf("%+v\n", *aciProfile.Name)
    }
  }
```

### LDAP

- [GetLDAPSetup](#getldapsetup)


```go
type LDAPSetup struct {
	Server                		*string  
	Port                   		*int64   
	BaseDN                 		*string  
	ServiceAccountDN       		*string  
	ServiceAccountPassword 		*string  
	StartTLS               		*bool    
	InsecureSkipVerify     		*bool    
}
```

### GetLDAPSetup

```go
func (s *Client) GetLDAPSetup() (*LDAPSetup, error)
```

##### Example
```go
  ldapSetup, err := client.GetLDAPSetup()
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%+v\n", *ldapSetup.Server)
  }
```

### RBAC

- [GetRole](#getrole)


```go
type Role struct {
	Role		 *string  
}
```

### GetRole

```go
func (s *Client) GetRole() (*Role, error)
```

##### Example
```go
  role, err := client.GetRole()
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%+v\n", *role.Role)
  }
```


DISCLAIMER:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
