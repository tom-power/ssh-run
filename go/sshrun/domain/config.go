package domain

import (
	"errors"
	"strings"

	. "github.com/tom-power/ssh-run/sshrun/utils/fp"
)

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts"`
	LocalhostIs           string `yaml:"localhostIs"`
}

func (c Config) HostNames() (string, error) {
	return c.hostNames(" "), nil
}

func (c Config) Host(hostName string) (Host, error) {
	host, err := Single(c.Hosts, isHostName(hostName))
	if err != nil {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + c.hostNames(", "))
	}
	return *host, nil
}

func (c Config) hostNames(sep string) string {
	names := Map(c.Hosts, toHostName)
	return strings.Join(names, sep)
}

var toHostName = func(host Host) string { return host.Name }

var isHostName = func(hostName string) func(host Host) bool {
	return func(host Host) bool {
		return host.Name == hostName
	}
}
