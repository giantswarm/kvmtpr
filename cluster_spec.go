package clusterspec

type ClusterSpec struct {
	Customer       string `json:"customer"`
	ClusterID      string `json:"clusterId"`
	WorkerReplicas int32  `json:"workerReplicas,omitempty"`

	K8sVmVersion string `json:"k8sVmVersion"`

	FlannelClientVersion string `json:"flannelClientVersion"`
	ClusterVNI           int32  `json:"clusterVNI,omitempty"`
	ClusterBackend       string `json:"clusterBackend"`
	ClusterNetwork       string `json:"clusterNetwork"`

	Machine Machine `json:"machine"`

	IngressController IngressController `json:"ingressController"`
}

type IngressController struct {
	KempVsIp              string `json:"kempVsIp"`
	KempVsName            string `json:"kempVsName"`
	KempVsPorts           string `json:"kempVsPorts"`
	KempVsSslAcceleration string `json:"kempVsSslAcceleration"`
	KempRsPort            string `json:"kempRsPort"`
	KempVsCheckPort       string `json:"kempVsCheckPort"`
	CloudflareIp          string `json:"cloudflareIp"`
	CloudflareDomain      string `json:"cloudflareDomain"`
}

type Machine struct {
	Registry string `json:"registry"`

	CalicoSubnet string `json:"calicoSubnet"`
	CalicoCidr   string `json:"calicoCidr"`
	K8sCalicoMtu string `json:"k8sCalicoMtu"`

	Certificates Certificates `json:"certificates"`

	MachineMem      string `json:"machineMem"`
	MachineCPUcores int32  `json:"machineCPUcores"`

	DockerExtraArgs     string `json:"dockerExtraArgs,omitempty"`
	NetworkSetupVersion string `json:"networkSetupVersion"`

	KubernetesConfig Kubernetes `json:"kubernetes"`
}

type Certificates struct {
	VaultToken        string `json:"vaulToken"`
	APIaltNames       string `json:"apiAltNames"`
	MasterServiceName string `json:"masterServiceName"`
}

type Kubernetes struct {
	Domain            string `json:"domain"`
	ETCDdomainName    string `json:"etcdDomainName"`
	MasterDomainName  string `json:"masterDomainName"`
	NodeLabels        string `json:"nodeLabels,omitempty"`
	ClusterIpRange    string `json:"clusterIpRange"`
	ClusterIpSubnet   string `json:"clusterIpSubnet"`
	MasterPort        string `json:"masterPort"`
	DnsIp             string `json:"dnsIp"`
	InsecurePort      string `json:"insecurePort"`
	SecurePort        string `json:"securePort"`
	WorkerServicePort string `json:"workerServicePort"`
}
