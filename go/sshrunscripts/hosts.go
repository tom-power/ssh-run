package sshrunscripts

import (
	"errors"
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
		hostsToHostName := func(host Host) string { return host.Name }
		return strings.Join(Map(hosts, hostsToHostName), " "), nil
	}
}

func GetHostsFromConfig(configBytes []byte) ([]Host, error) {
	hostsConfig := HostsConfig{}
	err := yaml.Unmarshal(configBytes, &hostsConfig)
	if err != nil {
		return nil, err
	}
	hosts := hostsConfig.Hosts
	if len(hosts) == 0 {
		return []Host{}, errors.New("no hosts found in config: " + string(configBytes))
	}
	return hosts, nil
}
