package flannel

type Flannel struct {
	// Backend is the Flannel backend type, e.g. vxlan.
	Backend string `json:"backend"`
	// Interface is the network interface name, e.g. bond0.3, or ens33.
	Interface string `json:"interface"`
	// Network is the subnet specification, e.g. 10.0.9.0/16.
	Network string `json:"network"`
	// Version is the Docker image tag.
	Version string `json:"version"`
	// VNI is the vxlan network identifier, e.g. 9.
	VNI int `json:"vni"`
}
