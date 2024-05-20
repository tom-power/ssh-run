package domain

import (
	"errors"
	"strings"

	"github.com/jucardi/go-streams/v2/streams"

	fp "github.com/rjNemo/underscore"
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
	host := streams.From[Host](c.Hosts).Filter(isHostName(hostName)).Distinct().First()
	if host.Ip == "" {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + c.hostNames(", "))
	}
	return host, nil
}

func (c Config) hostNames(sep string) string {
	names := fp.Map(c.Hosts, toHostName)
	return strings.Join(names, sep)
}

var toHostName = func(host Host) string { return host.Name }

var isHostName = func(hostName string) func(host Host) bool {
	return func(host Host) bool {
		return host.Name == hostName
	}
}
