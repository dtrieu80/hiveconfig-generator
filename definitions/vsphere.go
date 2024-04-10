package definitions

import (
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type VsphereHiveConfig struct {
	Metadata Metadata        `yaml:"metadata" json:"metadata"`
	Cluster  Cluster         `yaml:"cluster" json:"cluster"`
	Provider VsphereProvider `yaml:"provider" json:"provider"`
}

type VsphereProvider struct {
	Vsphere Vsphere `yaml:"vsphere" json:"vsphere"`
}

type Vsphere struct {
	ApiVIP            string             `yaml:"apiVIP" json:"apiVIP"`
	IngressVIP        string             `yaml:"ingressVIP" json:"ingressVIP"`
	ControlNodes      VirtControlNodes   `yaml:"controlNodes" json:"controlNodes"`
	ComputeNodes      []VirtComputeNodes `yaml:"computeNodes" json:"computeNodes"`
	ProviderSecretRef string             `yaml:"providerSecretRef" json:"providerSecretRef"`
}

type VirtControlNodes struct {
	Replicas       int `yaml:"replicas" json:"replicas"`
	Cpus           int `yaml:"cpus" json:"cpus"`
	CoresPerSocket int `yaml:"coresPerSocket" json:"coresPerSocket"`
	MemoryMB       int `yaml:"memoryMB" json:"memoryMB"`
	OsDiskSizeGB   int `yaml:"osDiskSizeGB" json:"osDiskSizeGB"`
}

type VirtComputeNodes struct {
	Id             int      `yaml:"id" json:"id"`
	Replicas       int      `yaml:"replicas" json:"replicas"`
	Cpus           int      `yaml:"cpus" json:"cpus"`
	CoresPerSocket int      `yaml:"coresPerSocket" json:"coresPerSocket"`
	MemoryMB       int      `yaml:"memoryMB" json:"memoryMB"`
	OsDiskSizeGB   int      `yaml:"osDiskSizeGB" json:"osDiskSizeGB"`
	Zones          []string `yaml:"zones" json:"zones"`
}

func (config VsphereHiveConfig) Default() error {
	// println("VsphereHiveConfig.Default 1")
	return nil
}

func (config VsphereHiveConfig) Validate() error {
	// println("VsphereHiveConfig.Validate 1")
	return nil
}

func (config VsphereHiveConfig) Filter(items []*yaml.RNode) ([]*yaml.RNode, error) {
	// println("VsphereHiveConfig.Filter 1")
	return items, nil
}
