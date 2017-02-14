package operator

type Operator struct {
	CertctlVersion      string `json:"certctl_ersion"`
	K8sVmVersion        string `json:"k8s_vm_version"`
	KubectlVersion      string `json:"kubectl_version"`
	NetworkSetupVersion string `json:"networkSetupVersion"`
}
