package config

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
)

type GetConfig = func() (domain.Config, error)

func (c ConfigFs) GetConfig() (domain.Config, error) {
	yamlConfig, err := c.getConfigFromYaml()
	if len(yamlConfig.Hosts) == 0 || yamlConfig.IncludeSshConfigHosts {
		sshConfigHosts, _ := c.getHostsFromSshConfig()
		yamlConfig.Hosts = append(yamlConfig.Hosts, sshConfigHosts...)
	}
	return yamlConfig, err
}
