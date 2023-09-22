package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/domain"
	"testing"
	"testing/fstest"
)

func Test_fileSysConfig(t *testing.T) {

	t.Run("get config fs", func(t *testing.T) {
		var testFsConfig = fstest.MapFS{
			"config.yml":                  {Data: []byte(configYamlText)},
		}

		sys := config.ConfigFs{
			Fsys:       testFsConfig,
			ConfigPath: "config.yml",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, expectedConfigFromYaml)
	})
}

var configYamlText = `
includeSshConfigHosts: true
localhostIs: testName1
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

var expectedConfigFromYaml = domain.Config{
	IncludeSshConfigHosts: true,
	LocalhostIs: "testName1",
	Hosts: []domain.Host{
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

func assertConfigEqual(t *testing.T, actual domain.Config, expected domain.Config) {
	if actual.IncludeSshConfigHosts != expected.IncludeSshConfigHosts {
		t.Errorf("actual.IncludeSshConfigHosts '%v' should equal '%v'", actual.IncludeSshConfigHosts, expected.IncludeSshConfigHosts)
	}
	if actual.LocalhostIs != expected.LocalhostIs {
		t.Errorf("actual.LocalhostIs '%v' should equal '%v'", actual.LocalhostIs, expected.LocalhostIs)
	}
	assertHostsEqual(t, actual.Hosts, expected.Hosts)
}

func assertHostsEqual(t *testing.T, actual []domain.Host, expected []domain.Host) {
	if len(actual) != len(expected) {
		t.Errorf("len(actual) '%v' should equal '%v'", len(actual), len(expected))
	}
	for i, host := range actual {
		if host != expected[i] {
			t.Errorf("actual host '%v' should equal '%v'", host, expected[i])
		}
	}
}
