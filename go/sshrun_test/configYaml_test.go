package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

var configText = `
sshOnNoCommand: true
hostsFromSshConfig: true
hosts:
  - host: 192.0.2.1
    user: testUser1
    name: testName1
    port: 22
  - host: 192.0.2.2
    user: testUser2
    name: testName2
    port: 23
    portTunnel: 1081
  - host: 192.0.2.3
    user: testUser3
    name: testName3
    port: 24`

func Test_configYaml(t *testing.T) {
	t.Run("can get config from yaml", func(t *testing.T) {
		actual, _ := config.GetConfigFromYamlBytes([]byte(configText))
		assertConfigEqual(t, actual, expectedConfigFromYaml)
	})
}

func assertConfigEqual(t *testing.T, actual shared.Config, expected shared.Config) {
	if actual.SshOnNoCommand != expected.SshOnNoCommand {
		t.Errorf("'%v' should equal '%v'", actual.SshOnNoCommand, expected.SshOnNoCommand)
	}
	if actual.HostsFromSshConfig != expected.HostsFromSshConfig {
		t.Errorf("'%v' should equal '%v'", actual.HostsFromSshConfig, expected.HostsFromSshConfig)
	}
	if actual.Hosts[0] != expected.Hosts[0] {
		t.Errorf("'%v' should equal '%v'", actual.Hosts[0], expected.Hosts[0])
	}
}

var expectedConfigFromYaml = shared.Config{
	SshOnNoCommand:     true,
	HostsFromSshConfig: true,
	Hosts: []shared.Host{
		{
			Name: "testName1",
			User: "testUser1",
			Host: "192.0.2.1",
			Port: "22",
		},
		{
			Name:       "testName2",
			User:       "testUser2",
			Host:       "192.0.2.2",
			Port:       "23",
			PortTunnel: "1081",
		},
		{
			Name: "testName3",
			User: "testUser3",
			Host: "192.0.2.3",
			Port: "24",
		},
	},
}
