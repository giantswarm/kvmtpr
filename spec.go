package kvmtpr

import (
	"github.com/giantswarm/kvmtpr/spec/calico"
	"github.com/giantswarm/kvmtpr/spec/cloudflare"
	"github.com/giantswarm/kvmtpr/spec/cluster"
	"github.com/giantswarm/kvmtpr/spec/customer"
	"github.com/giantswarm/kvmtpr/spec/docker"
	"github.com/giantswarm/kvmtpr/spec/etcd"
	"github.com/giantswarm/kvmtpr/spec/flannel"
	"github.com/giantswarm/kvmtpr/spec/kubernetes"
	"github.com/giantswarm/kvmtpr/spec/node"
	"github.com/giantswarm/kvmtpr/spec/operator"
	"github.com/giantswarm/kvmtpr/spec/vault"
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
