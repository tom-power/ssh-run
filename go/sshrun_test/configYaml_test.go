package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
	"testing/fstest"
)

var configText = `
includeSshConfigHosts: false
checkRemoteForScripts: true
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
		testFs := fstest.MapFS{
			"config.yaml": {
				Data: []byte(configText),
			},
		}
		actual, err := config.GetConfigFromYaml("config.yaml", testFs)
		if err != nil {
			t.Errorf(err.Error())
		}
		assertConfigEqual(t, actual, expectedConfigFromYaml)
	})
}

func assertConfigEqual(t *testing.T, actual shared.Config, expected shared.Config) {
	if actual.IncludeSshConfigHosts != expected.IncludeSshConfigHosts {
		t.Errorf("actual.IncludeSshConfigHosts '%v' should equal '%v'", actual.IncludeSshConfigHosts, expected.IncludeSshConfigHosts)
	}
	if actual.CheckRemoteForScripts != expected.CheckRemoteForScripts {
		t.Errorf("actual.CheckRemoteForScripts '%v' should equal '%v'", actual.CheckRemoteForScripts, expected.CheckRemoteForScripts)
	}
	assertHostsEqual(t, actual.Hosts, expected.Hosts)
}

var expectedConfigFromYaml = shared.Config{
	IncludeSshConfigHosts: false,
	CheckRemoteForScripts: true,
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
