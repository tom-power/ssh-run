package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetHosts = func() (string, error)

var GetHostsFromConfig = func(config shared.Config) GetHosts {
	return func() (string, error) {
		return shared.HostsToHostNames(config.Hosts, " "), nil
	}
}
