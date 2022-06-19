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

func (config Config) HostNames() (string, error) {
	return config.hostNames(" "), nil
}

func (config Config) Host(hostName string) (Host, error) {
	host, err := shared.Single(config.Hosts, func(host Host) bool { return host.Name == hostName })
	if err != nil {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + config.hostNames(", "))
	}
	return *host, nil
}

func (config Config) hostNames(sep string) string {
	names := shared.Map(config.Hosts, toName)
	return strings.Join(names, sep)
}

func toName(host Host) string { return host.Name }
