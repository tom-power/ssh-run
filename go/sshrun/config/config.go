package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetConfig = func() (shared.Config, error)

var GetConfigFromFileSystem = func(
	getConfigFromYaml GetConfigFrom,
	getHostsFromSshConfig GetHostsFrom,
) (shared.Config, error) {
	yaml, err := getConfigFromYaml()
	if err != nil {
		return shared.Config{}, err
	}
	if yaml.HostsFromSshConfig {
		sshConfigHosts, err := GetSshConfigHosts()
		if err != nil {
			return shared.Config{}, err
		}
		return shared.Config{
			Hosts:              append(yaml.Hosts, sshConfigHosts...),
			SshOnNoCommand:     yaml.SshOnNoCommand,
			HostsFromSshConfig: yaml.HostsFromSshConfig,
		}, nil
	}
	return yaml, nil
}
