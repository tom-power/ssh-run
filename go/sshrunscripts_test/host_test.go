package sshrunscripts

import (
	"testing"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

var configText = `
hosts:
  - name: testName1
    user: testUser1
    ip: 192.0.2.1
    portSsh: 22
  - name: testName2
    user: testUser2
    ip: 192.0.2.2
    portSsh: 23
    portTunnel: 1081
  - name: testName3
    user: testUser3
    ip: 192.0.2.3
    portSsh: 24`

var testGetHostFromConf = sshrunscripts.GetHostFromConf([]byte(configText))

func Test_host(t *testing.T) {
	t.Run("can get host from conf", func(t *testing.T) {
		host, _ := testGetHostFromConf("testName2", "")

		var expectedHost = sshrunscripts.Host{
			Name:       "testName2",
			User:       "testUser2",
			Ip:         "192.0.2.2",
			PortSsh:    "23",
			PortTunnel: "1081",
		}
		if host != expectedHost {
			t.Errorf("'%v' should equal '%v'", host, expectedHost)
		}
	})
}
