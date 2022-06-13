package sshrun

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetHost = func(hostName string) (shared.Host, error)

var GetHostFromConfig = func(config shared.Config) GetHost {
	return func(hostName string) (shared.Host, error) {
		return getHost(hostName, config.Hosts)
	}
}

func getHost(hostName string, hosts []shared.Host) (shared.Host, error) {
	var hasHostName = func(host shared.Host) bool { return host.Name == hostName }
	host, err := shared.Single(hosts, hasHostName)
	if err != nil {
		return shared.Host{}, getError(hostName, hosts)
	}
	return *host, nil
}

func getError(hostName string, hosts []shared.Host) error {
	return errors.New("couldn't find host " + hostName + " in " + shared.HostsToHostNames(hosts, ", "))
}
