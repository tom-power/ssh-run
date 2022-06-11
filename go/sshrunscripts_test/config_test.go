package sshrunscripts

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/config"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"testing"
)

var configText = `
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

var testGetConfig = config.GetConfigFromYaml([]byte(configText))

func Test_config(t *testing.T) {
	t.Run("can get host from conf", func(t *testing.T) {
		config, _ := testGetConfig()

		var expectedConfig = shared.Config{
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

		if config.Hosts[0] != expectedConfig.Hosts[0] {
			t.Errorf("'%v' should equal '%v'", config, expectedConfig)
		}
	})
}
