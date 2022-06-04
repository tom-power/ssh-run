package sshrunscripts

import (
	"errors"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"gopkg.in/yaml.v2"
	"strings"
)

type GetHosts = func() (string, error)

var GetHostsFromConf = func(configBytes []byte) GetHosts {
	return func() (string, error) {
		hosts, err := GetHostsFromConfig(configBytes)
		if err != nil {
			return "", err
		}
		hostsToHostName := func(host shared.Host) string { return host.Name }
		return strings.Join(shared.Map(hosts, hostsToHostName), " "), nil
	}
}

func GetHostsFromConfig(configBytes []byte) ([]shared.Host, error) {
	hostsConfig := HostsConfig{}
	err := yaml.Unmarshal(configBytes, &hostsConfig)
	if err != nil {
		return nil, err
	}
	hosts := hostsConfig.Hosts
	if len(hosts) == 0 {
		return []shared.Host{}, errors.New("no hosts found in config: " + string(configBytes))
	}
	return hosts, nil
}

type HostsConfig struct {
	Hosts []shared.Host `yaml:"hosts"`
}
