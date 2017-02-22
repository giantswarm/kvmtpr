package etcd

type Etcd struct {
	Domain string `json:"domain" yaml:"domain"`
	IP     string `json:"ip" yaml:"ip"`
	Port   string `json:"port" yaml:"port"`
	Prefix string `json:"prefix" yaml:"prefix"`
}
