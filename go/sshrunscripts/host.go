package sshrunscripts

import (
	"errors"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
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
