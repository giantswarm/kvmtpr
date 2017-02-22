package node

type Node struct {
	Memory string `json:"memory" yaml:"memory"`
	CPUs   int    `json:"cpus" yaml:"cpus"`
}
