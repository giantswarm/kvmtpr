package clusterspec

type ClusterSpec struct {
	Customer       string `json:"customer"`
	ClusterID      string `json:"clusterId"`
	WorkerReplicas int32  `json:"workerReplicas,omitempty"`

	K8sVmVersion         string `json:"k8sVmVersion"`
	FlannelClientVersion string `json:"flannelClientVersion"`

	ClusterVNI     int32  `json:"clusterVNI,omitempty"`
	ClusterBackend string `json:"clusterBackend"`
	ClusterNetwork string `json:"clusterNetwork"`

	KubernetesConfig KubernetesConfiguration `json:"kubernetesConfiguration"`

	MachineMem      string `json:"machineMem"`
	MachineCPUcores int32  `json:"machineCPUcores"`

	VaultToken string `json:"vaulToken"`

	// Ingress Controller
	KempVsIp              string `json:"kempVsIp"`
	KempVsName            string `json:"kempVsName"`
	KempVsPorts           string `json:"kempVsPorts"`
	KempVsSslAcceleration string `json:"kempVsSslAcceleration"`
	KempRsPort            string `json:"kempRsPort"`
	KempVsCheckPort       string `json:"kempVsCheckPort"`
	CloudflareIp          string `json:"cloudflareIp"`
	CloudflareDomain      string `json:"cloudflareDomain"`
}

type KubernetesConfiguration struct {
	Registry string `json:"registry"`

	CalicoSubnet string `json:"calicoSubner"`
	CalicoCidr   string `json:"calicoCidr"`
	K8sCalicoMtu string `json:"k8sCalicoMtu"`

	DockerExtraArgs     string `json:"dockerExtraArgs,omitempty"`
	NetworkSetupVersion string `json:"networkSetupVersion"`

	APIaltNames       string `json:"apiAltNames"`
	Domain            string `json:"domain"`
	ETCDdomainName    string `json:"etcdDomainName"`
	MasterDomainName  string `json:"masterDomainName"`
	MasterServiceName string `json:"masterServiceName"`
	NodeLabels        string `json:"nodeLabels,omitempty"`
	ClusterIpRange    string `json:"clusterIpRange"`
	ClusterIpSubnet   string `json:"clusterIpSubnet"`
	MasterPort        string `json:"masterPort"`
	DnsIp             string `json:"dnsIp"`
	InsecurePort      string `json:"insecurePort"`
	SecurePort        string `json:"securePort"`
	WorkerServicePort string `json:"workerServicePort"`
}
