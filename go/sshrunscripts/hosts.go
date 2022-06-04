package sshrunscripts

import (
	"errors"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"gopkg.in/yaml.v2"
	"strings"
)

type GetHost = func(hostName string, localUserName string) (shared.Host, error)

var GetHostFromConf = func(configBytes []byte) GetHost {
	return func(hostName string, localUserName string) (shared.Host, error) {
		hosts, err := GetHostsFromConfig(configBytes)
		if err != nil {
			return shared.Host{}, err
		}
		return getHost(hostName, hosts)
	}
}

type HostsConfig struct {
	Hosts []shared.Host `yaml:"hosts"`
}

func getHost(hostName string, hosts []shared.Host) (shared.Host, error) {
	for i := range hosts {
		host := hosts[i]
		if host.Name == hostName {
			return host, nil
		}
	}
	return shared.Host{}, errors.New("couldn't find host " + hostName + " in " + strings.Join(names(hosts), ", "))
}

func names(hosts []shared.Host) []string {
	var names []string
	for i := range hosts {
		names = append(names, hosts[i].Name)
	}
	return names
}

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
