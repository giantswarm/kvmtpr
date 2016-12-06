package clusterspec

type ClusterSpec struct {
	Customer  string `json:"customer"`
	ClusterID string `json:"clusterId"`

	K8sVmVersion string `json:"k8sVmVersion"`

	FlannelConfiguration FlannelConfiguration `json::"flannelConfiguration"`

	Certificates Certificates `json:"certificates"`

	Worker Worker `json:"worker"`
	Master Master `json:"master"`

	IngressController IngressController `json:"ingressController"`
}

type FlannelConfiguration struct {
	Version        string `json:"version"`
	ClusterVNI     int32  `json:"clusterVNI,omitempty"`
	ClusterBackend string `json:"clusterBackend"`
	ClusterNetwork string `json:"clusterNetwork"`
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

type Certificates struct {
	VaultToken        string `json:"vaulToken"`
	APIaltNames       string `json:"apiAltNames"`
	MasterServiceName string `json:"masterServiceName"`
}

type Machine struct {
	Registry            string       `json:"registry"`
	Capabilities        Capabilities `json:"capabilities"`
	NetworkSetupVersion string       `json:"networkSetupVersion"`
	KubernetesConfig    Kubernetes   `json:"kubernetes"`
}

type Capabilities struct {
	Memory   string `json:"memory"`
	CPUcores int32  `json:"cpuCores"`
}

type Worker struct {
	Machine
	Replicas int32 `json:"workerReplicas,omitempty"`

	DockerExtraArgs   string `json:"dockerExtraArgs,omitempty"`
	K8sCalicoMtu      string `json:"k8sCalicoMtu"`
	NodeLabels        string `json:"nodeLabels,omitempty"`
	WorkerServicePort string `json:"workerServicePort"`
}

type Master struct {
	Machine

	CalicoSubnet string `json:"calicoSubnet"`
	CalicoCidr   string `json:"calicoCidr"`
}

type Kubernetes struct {
	Domain           string `json:"domain"`
	ETCDdomainName   string `json:"etcdDomainName"`
	MasterDomainName string `json:"masterDomainName"`
	ClusterIpRange   string `json:"clusterIpRange"`
	ClusterIpSubnet  string `json:"clusterIpSubnet"`
	MasterPort       string `json:"masterPort"`
	DnsIp            string `json:"dnsIp"`
	InsecurePort     string `json:"insecurePort"`
	SecurePort       string `json:"securePort"`
}
