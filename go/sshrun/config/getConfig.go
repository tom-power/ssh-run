package config

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
)

type GetConfig = func() (domain.Config, error)

func (c ConfigFs) GetConfig() (domain.Config, error) {
	yamlConfig, err := c.getConfigFromYaml()
	if err != nil || len(yamlConfig.Hosts) == 0 || yamlConfig.IncludeSshConfigHosts == true {
		sshConfigHosts, _ := c.getHostsFromSshConfig()
		return withHosts(yamlConfig, sshConfigHosts), nil
	}
	return yamlConfig, nil
}

func withHosts(config domain.Config, hosts []domain.Host) domain.Config {
	return domain.Config{
		Hosts:                 append(config.Hosts, hosts...),
		IncludeSshConfigHosts: config.IncludeSshConfigHosts,
	}
}
