package definitions

import (
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/parser"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type AzureHiveConfig struct {
	ApiVersion string        `yaml:"apiVersion" json:"apiVersion"`
	Kind       string        `yaml:"kind" json:"kind"`
	Metadata   Metadata      `yaml:"metadata" json:"metadata"`
	Cluster    Cluster       `yaml:"cluster" json:"cluster"`
	Provider   AzureProvider `yaml:"provider" json:"provider"`
}

type AzureProvider struct {
	Azure Azure `yaml:"azure" json:"azure"`
}

type Azure struct {
	ResourceGroup string              `yaml:"resourceGroup" json:"resourceGroup"`
	Region        string              `yaml:"region" json:"region"`
	ControlNodes  CloudControlNodes   `yaml:"controlNodes" json:"controlNodes"`
	ComputeNodes  []CloudComputeNodes `yaml:"computeNodes" json:"computeNodes"`
}

type CloudControlNodes struct {
	Replicas     int    `yaml:"replicas" json:"replicas"`
	Size         string `yaml:"size" json:"size"`
	OsDiskSizeGB int    `yaml:"osDiskSizeGB" json:"osDiskSizeGB"`
}

type CloudComputeNodes struct {
	Id           int      `yaml:"id" json:"id"`
	Replicas     int      `yaml:"replicas" json:"replicas"`
	Size         string   `yaml:"size" json:"size"`
	OsDiskSizeGB int      `yaml:"osDiskSizeGB" json:"osDiskSizeGB"`
	Zones        []string `yaml:"zones" json:"zones"`
}

func (config AzureHiveConfig) Default() error {
	// println("AzureHiveConfig.Default 1")
	return nil
}

func (config AzureHiveConfig) Validate() error {
	// println("AzureHiveConfig.Validate 1")
	return nil
}

func (config AzureHiveConfig) Filter(items []*yaml.RNode) ([]*yaml.RNode, error) {
	// println("AzureHiveConfig.Filter 1")
	filter := framework.TemplateProcessor{
		ResourceTemplates: []framework.ResourceTemplate{{
			TemplateData: &config,
			Templates:    parser.TemplateFiles("templates/azure"),
		}},
	}
	return filter.Filter(items)
}
