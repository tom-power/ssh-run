package domain

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"strings"
)

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts"`
	CheckRemoteForScripts bool   `yaml:"checkRemoteForScripts"`
}

func (config Config) HostNames() (string, error) {
	return config.hostNames(" "), nil
}

func (config Config) Host(hostName string) (Host, error) {
	var hasHostName = func(host Host) bool { return host.Name == hostName }
	hostFromHosts, err := shared.Single(config.Hosts, hasHostName)
	if err != nil {
		return Host{}, errors.New("couldn't find host " + hostName + " in " + config.hostNames(", "))
	}
	return hostWithRemote(config, *hostFromHosts), nil
}

func (config Config) hostNames(sep string) string {
	names := shared.Map(config.Hosts, toName)
	return strings.Join(names, sep)
}

func hostWithRemote(config Config, hostOther Host) Host {
	return Host{
		Host:        hostOther.Host,
		User:        hostOther.User,
		Name:        hostOther.Name,
		Port:        hostOther.Port,
		PortTunnel:  hostOther.PortTunnel,
		CheckRemote: config.CheckRemoteForScripts,
	}
}

func toName(host Host) string { return host.Name }
