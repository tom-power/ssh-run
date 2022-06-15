package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
	"testing/fstest"
)

var hostsFromSshConfigText = `
Host *
  ServerAliveInterval 60
  IgnoreUnknown PortTunnel

Host 192.168.6*
  StrictHostKeyChecking no

Host localhost
  User test
  Hostname localhost

Host github.com
  User git
  Hostname github.com
  PreferredAuthentications publickey
  IdentityFile ~/.ssh/id_rsa`

func Test_hostsSshConfig(t *testing.T) {
	t.Run("can get config from yaml", func(t *testing.T) {
		testFs := fstest.MapFS{
			".ssh/config": {
				Data: []byte(hostsFromSshConfigText),
			},
		}
		actual, err := config.GetHostsFromSshConfig(".ssh/config", testFs)
		if err != nil {
			t.Errorf(err.Error())
		}
		assertHostsEqual(t, actual, expectedHostsFromSshConfig)
	})
}

func assertHostsEqual(t *testing.T, actual []shared.Host, expected []shared.Host) {
	if len(actual) != len(expected) {
		t.Errorf("len(actual) '%v' should equal '%v'", len(actual), len(expected))
	}
	for i, host := range actual {
		if host != expected[i] {
			t.Errorf("actual host '%v' should equal '%v'", host, expected[i])
		}
	}
}

var expectedHostsFromSshConfig = []shared.Host{
	{
		Name: "localhost",
		User: "test",
		Host: "localhost",
		Port: "22",
	},
	{
		Name: "github.com",
		User: "git",
		Host: "github.com",
		Port: "22",
	},
}
