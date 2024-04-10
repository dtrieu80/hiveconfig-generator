package definitions

import (
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/parser"
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
	Vcenter          string             `yaml:"vcenter" json:"vcenter"`
	Username         string             `yaml:"username" json:"username"`
	Password         string             `yaml:"password" json:"password"`
	Datacenter       string             `yaml:"datacenter" json:"datacenter"`
	Cluster          string             `yaml:"cluster" json:"cluster"`
	Network          string             `yaml:"network" json:"network"`
	DefaultDatastore string             `yaml:"defaultDatastore" json:"defaultDatastore"`
	Folder           string             `yaml:"folder" json:"folder"`
	ApiVIP           string             `yaml:"apiVIP" json:"apiVIP"`
	IngressVIP       string             `yaml:"ingressVIP" json:"ingressVIP"`
	ControlNodes     VirtControlNodes   `yaml:"controlNodes" json:"controlNodes"`
	ComputeNodes     []VirtComputeNodes `yaml:"computeNodes" json:"computeNodes"`
	//ProviderSecretRef string             `yaml:"providerSecretRef" json:"providerSecretRef"`
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
	filter := framework.TemplateProcessor{
		ResourceTemplates: []framework.ResourceTemplate{{
			TemplateData: &config,
			Templates:    parser.TemplateFiles("templates/vsphere"),
		}},
	}
	return filter.Filter(items)
}
