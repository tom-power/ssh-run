package shared

import (
	"errors"
	"fmt"
	"strings"
)

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts"`
	CheckRemoteForScripts bool   `yaml:"checkRemoteForScripts"`
}

func (config Config) HostNames() (string, error) {
	return HostsToHostNames(config.Hosts, " "), nil
}

func (config Config) Host(hostName string) (Host, error) {
	var hasHostName = func(host Host) bool { return host.Name == hostName }
	host, err := Single(config.Hosts, hasHostName)
	if err != nil {
		return Host{}, getError(hostName, config.Hosts)
	}
	return *host, nil
}

func getError(hostName string, hosts []Host) error {
	return errors.New("couldn't find host " + hostName + " in " + HostsToHostNames(hosts, ", "))
}

type Host struct {
	Host       string
	User       string
	Name       string
	Port       string
	PortTunnel string `yaml:"portTunnel"`
}

func (h Host) ToString() string {
	return strings.ReplaceAll(fmt.Sprintf("%#v", h), "shared.", "")
}
