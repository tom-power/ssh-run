package domain

import (
	"errors"
	"strings"

	"github.com/tom-power/ssh-run/sshrun/shared/generic"
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
	host, err := generic.Single(c.Hosts, func(host Host) bool { return host.Name == hostName })
	if err != nil {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + c.hostNames(", "))
	}
	return *host, nil
}

func (c Config) hostNames(sep string) string {
	names := generic.Map(c.Hosts, toName)
	return strings.Join(names, sep)
}

func toName(host Host) string { return host.Name }
