package definitions

type Metadata struct {
	Name string `yaml:"name" json:"name"`
}

type Cluster struct {
	BaseDomain          string            `yaml:"baseDomain" json:"baseDomain"`
	OcpVersionInstalled string            `yaml:"ocpVersionInstalled" json:"ocpVersionInstalled"`
	OcpVersionDesired   string            `yaml:"ocpVersionDesired" json:"ocpVersionDesired"`
	Networking          Networking        `yaml:"networking" json:"networking"`
	ClusterPowerState   string            `yaml:"clusterPowerState" json:"clusterPowerState"`
	AcmClusterSetName   string            `yaml:"acmClusterSetName" json:"acmClusterSetName"`
	AcmClusterLabels    map[string]string `yaml:"acmClusterLabels" json:"acmClusterLabels"`
	SshPubkey           string            `yaml:"sshPubkey" json:"sshPubkey"`
}

type Networking struct {
	Type           string `yaml:"type" json:"type"`
	ClusterNetwork string `yaml:"clusterNetwork" json:"clusterNetwork"`
	HostPrefix     string `yaml:"hostPrefix" json:"hostPrefix"`
	MachineNetwork string `yaml:"machineNetwork" json:"machineNetwork"`
	ServiceNetwork string `yaml:"serviceNetwork" json:"serviceNetwork"`
}
