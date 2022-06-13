package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetHostsFrom = func() ([]shared.Host, error)

func GetSshConfigHosts() ([]shared.Host, error) {
	return []shared.Host{}, nil
}
