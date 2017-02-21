package operator

type Operator struct {
	CertctlVersion      string `json:"certctlVersion"`
	K8sVmVersion        string `json:"k8sVmVersion"`
	KubectlVersion      string `json:"kubectlVersion"`
	NetworkSetupVersion string `json:"networkSetupVersion"`
}
