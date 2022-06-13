package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetConfig = func() (shared.Config, error)

var GetConfigFromFileSystem = func(
	getConfigFromYaml GetConfigFrom,
	getHostsFromSshConfig GetHostsFrom,
) (shared.Config, error) {
	yamlConfig, err := getConfigFromYaml()
	if err != nil {
		return shared.Config{}, err
	}
	if yamlConfig.HostsFromSshConfig {
		sshConfigHosts, err := GetSshConfigHosts()
		if err != nil {
			return shared.Config{}, err
		}
		config := shared.Config{
			Hosts:              append(yamlConfig.Hosts, sshConfigHosts...),
			HostsFromSshConfig: yamlConfig.HostsFromSshConfig,
		}
		return config, nil
	}
	return yamlConfig, nil
}
