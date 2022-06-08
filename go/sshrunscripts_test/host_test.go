package sshrunscripts

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"testing"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
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

var testGetHostFromConf = sshrunscripts.GetHostFromConf([]byte(configText))

func Test_host(t *testing.T) {
	t.Run("can get host from conf", func(t *testing.T) {
		host, _ := testGetHostFromConf("testName2", "")

		var expectedHost = shared.Host{
			Name:       "testName2",
			User:       "testUser2",
			Host:       "192.0.2.2",
			Port:       "23",
			PortTunnel: "1081",
		}
		if host != expectedHost {
			t.Errorf("'%v' should equal '%v'", host, expectedHost)
		}
	})
}
