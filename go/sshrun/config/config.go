package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetConfig = func() (shared.Config, error)

var GetConfigFromFileSystem = func(
	getHostsFromSshConfig GetHostsFrom,
	getConfigFromYaml GetConfigFrom,
) (shared.Config, error) {
	yamlConfig, err := getConfigFromYaml()
	if err != nil {
		return shared.Config{}, err
	}
	if yamlConfig.IncludeSshConfig == true || len(yamlConfig.Hosts) == 0 {
		sshConfigHosts, err := getHostsFromSshConfig()
		if err != nil {
			return shared.Config{}, err
		}
		config := shared.Config{
			Hosts:            append(yamlConfig.Hosts, sshConfigHosts...),
			IncludeSshConfig: yamlConfig.IncludeSshConfig,
		}
		return config, nil
	}
	return yamlConfig, nil
}
