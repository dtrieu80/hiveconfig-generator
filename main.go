package main

import (
	"hiveconfig-generator/definitions"
	"os"

	"github.com/pkg/errors"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/kio/filters"
)

func processHiveProvider(rl *framework.ResourceList) error {
	// List of HiveConfigs per Provider
	p := framework.VersionedAPIProcessor{FilterProvider: framework.GVKFilterMap{
		"AzureHiveConfig": map[string]kio.Filter{
			"fleetmgmt.verse.dyt/v1alpha1": &definitions.AzureHiveConfig{},
		},
		"VsphereHiveConfig": map[string]kio.Filter{
			"fleetmgmt.verse.dyt/v1alpha1": &definitions.VsphereHiveConfig{},
		},
	}}

	// Process functionConfig with correct Provider
	if err := p.Process(rl); err != nil {
		return errors.Wrap(err, "Processing Hive Provider")
	}

	var err error

	rl.Items, err = filters.FormatFilter{UseSchema: true}.Filter(rl.Items)
	if err != nil {
		return errors.Wrap(err, "Formatting Output")
	}

	return nil
}

func main() {
	cmd := command.Build(framework.ResourceListProcessorFunc(processHiveProvider), command.StandaloneDisabled, false)
	//command.AddGenerateDockerfile(cmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
