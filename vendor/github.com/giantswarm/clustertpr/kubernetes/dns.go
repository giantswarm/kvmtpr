package kubernetes

type DNS struct {
	// IP is the kube DNS cluster IP, e.g. 172.31.0.10.
	IP string `json:"ip" yaml:"ip"`
}
