package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type Config = func() (shared.Config, error)

func (fsys FileSys) Config() (shared.Config, error) {
	yamlConfig, err := fsys.getConfigFromYaml()
	if err != nil || len(yamlConfig.Hosts) == 0 || yamlConfig.IncludeSshConfigHosts == true {
		sshConfigHosts, _ := fsys.getHostsFromSshConfig()
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
