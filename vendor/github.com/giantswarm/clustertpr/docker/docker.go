package docker

type Docker struct {
	Registry Registry `json:"registry" yaml:"registry"`
	Daemon   Daemon   `json:"daemon" yaml:"daemon"`
}
