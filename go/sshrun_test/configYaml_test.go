package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"strings"
	"testing"
)

var configText = `
includeSshConfig: true
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
		actual, err := config.GetConfigFromYamlReader(strings.NewReader(configText))
		if err != nil {
			t.Errorf(err.Error())
		}
		assertConfigEqual(t, actual, expectedConfigFromYaml)
	})
}

func assertConfigEqual(t *testing.T, actual shared.Config, expected shared.Config) {
	if actual.IncludeSshConfig != expected.IncludeSshConfig {
		t.Errorf("actual.IncludeSshConfig '%v' should equal '%v'", actual.IncludeSshConfig, expected.IncludeSshConfig)
	}
	if len(actual.Hosts) != len(expected.Hosts) {
		t.Errorf("len(actual.Hosts) '%v' should equal '%v'", len(actual.Hosts), len(expected.Hosts))
	}
	for i, host := range actual.Hosts {
		if host != expected.Hosts[i] {
			t.Errorf("actual host '%v' should equal '%v'", host, expected.Hosts[i])
		}
	}
}

var expectedConfigFromYaml = shared.Config{
	IncludeSshConfig: true,
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
