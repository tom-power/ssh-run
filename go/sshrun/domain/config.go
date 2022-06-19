package domain

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"strings"
)

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts"`
}

func (c Config) HostNames() (string, error) {
	return c.hostNames(" "), nil
}

func (c Config) Host(hostName string) (Host, error) {
	host, err := shared.Single(c.Hosts, func(host Host) bool { return host.Name == hostName })
	if err != nil {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + c.hostNames(", "))
	}
	return *host, nil
}

func (c Config) hostNames(sep string) string {
	names := shared.Map(c.Hosts, toName)
	return strings.Join(names, sep)
}

func toName(host Host) string { return host.Name }
