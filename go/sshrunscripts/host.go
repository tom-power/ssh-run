package sshrunscripts

import (
	"errors"
	"strings"
)

type Host struct {
	Name       string
	User       string
	Ip         string
	PortSsh    string `yaml:"portSsh"`
	PortTunnel string `yaml:"portTunnel"`
}

type GetHost = func(hostName string, localUserName string) (Host, error)

var GetHostFromConf = func(configBytes []byte) GetHost {
	return func(hostName string, localUserName string) (Host, error) {
		hosts, err := GetHostsFromConfig(configBytes)
		if err != nil {
			return Host{}, err
		}
		return getHost(hostName, hosts)
	}
}

type HostsConfig struct {
	Hosts []Host `yaml:"hosts"`
}

func getHost(hostName string, hosts []Host) (Host, error) {
	for i := range hosts {
		host := hosts[i]
		if host.Name == hostName {
			return host, nil
		}
	}
	return Host{}, errors.New("couldn't find host " + hostName + " in " + strings.Join(names(hosts), ", "))
}

func names(hosts []Host) []string {
	var names []string
	for i := range hosts {
		names = append(names, hosts[i].Name)
	}
	return names
}
