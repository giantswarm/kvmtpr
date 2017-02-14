package kvmtpr

import (
	"github.com/giantswarm/clusterspec/spec/calico"
	"github.com/giantswarm/clusterspec/spec/cloudflare"
	"github.com/giantswarm/clusterspec/spec/cluster"
	"github.com/giantswarm/clusterspec/spec/customer"
	"github.com/giantswarm/clusterspec/spec/docker"
	"github.com/giantswarm/clusterspec/spec/etcd"
	"github.com/giantswarm/clusterspec/spec/flannel"
	"github.com/giantswarm/clusterspec/spec/kubernetes"
	"github.com/giantswarm/clusterspec/spec/node"
	"github.com/giantswarm/clusterspec/spec/operator"
	"github.com/giantswarm/clusterspec/spec/vault"
)

type Spec struct {
	Calico     calico.Calico         `json:"calico"`
	Cloudflare cloudflare.Cloudflare `json:"cloudflare"`
	Cluster    cluster.Cluster       `json:"cluster"`
	Customer   customer.Customer     `json:"customer"`
	Docker     docker.Docker         `json:"docker"`
	Etcd       etcd.Etcd             `json:"etcd"`
	Flannel    flannel.Flannel       `json:"flannel"`
	Kubernetes kubernetes.Kubernetes `json:"kubernetes"`
	Masters    []node.Node           `json:"masters"`
	Operator   operator.Operator     `json:"operator"`
	Vault      vault.Vault           `json:"vault"`
	Workers    []node.Node           `json:"workers"`
}
