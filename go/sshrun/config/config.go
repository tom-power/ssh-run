package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func GetConfigUsing(
	hostsFromSshConfig GetHosts,
	configFromYaml GetConfig,
) (shared.Config, error) {
	yamlConfig, err := configFromYaml()
	if err != nil {
		return shared.Config{}, err
	}
	if yamlConfig.IncludeSshConfigHosts == true || len(yamlConfig.Hosts) == 0 {
		sshConfigHosts, err := hostsFromSshConfig()
		if err != nil {
			return shared.Config{}, err
		}
		return withHosts(yamlConfig, sshConfigHosts), nil
	}
	return yamlConfig, nil
}

func withHosts(config shared.Config, hosts []shared.Host) shared.Config {
	return shared.Config{
		Hosts:                 append(config.Hosts, hosts...),
		IncludeSshConfigHosts: config.IncludeSshConfigHosts,
		CheckRemoteForScripts: config.CheckRemoteForScripts,
	}
}
