package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/domain"
	"testing"
	"testing/fstest"
)

func Test_fileSysConfig(t *testing.T) {

	t.Run("get hosts from ssh config when IncludeSshConfigHosts fs", func(t *testing.T) {
		sys := config.ConfigFs{
			Fsys:       testFsConfig,
			ConfigPath: "config.yaml",
			SshPath:    ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{
			Hosts:                 append(expectedConfigFromYaml.Hosts, sshConfigHosts...),
			IncludeSshConfigHosts: true,
		})
	})
	t.Run("don't get hosts from ssh config when not IncludeSshConfigHosts fs", func(t *testing.T) {
		sys := config.ConfigFs{
			Fsys:       testFsConfig,
			ConfigPath: "configNoHostsFromConfig.yaml",
			SshPath:    ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{
			Hosts:                 append(expectedConfigFromYaml.Hosts),
			IncludeSshConfigHosts: false,
		})
	})
	t.Run("get hosts from ssh config when nothing in yaml", func(t *testing.T) {
		sys := config.ConfigFs{
			Fsys:       testFsConfig,
			ConfigPath: "blahBlah.yaml",
			SshPath:    ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{Hosts: sshConfigHosts})
	})
}

var testFsConfig = fstest.MapFS{
	".ssh/config":                  {Data: []byte(hostsSshConfigText)},
	"config.yaml":                  {Data: []byte(configYamlText)},
	"configNoHostsFromConfig.yaml": {Data: []byte(configYamlNoHostsFromSshText)},
}

var hostsSshConfigText = `
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

var configYamlText = `
includeSshConfigHosts: true
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

var configYamlNoHostsFromSshText = `
includeSshConfigHosts: false
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
	IncludeSshConfigHosts: false,
	CheckRemoteForScripts: true,
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

var sshConfigHosts = []domain.Host{{
	Name: "localhost",
	User: "test",
	Host: "localhost",
	Port: "22",
}, {
	Name: "github.com",
	User: "git",
	Host: "github.com",
	Port: "22",
}}

func assertConfigEqual(t *testing.T, actual domain.Config, expected domain.Config) {
	if actual.IncludeSshConfigHosts != expected.IncludeSshConfigHosts {
		t.Errorf("actual.IncludeSshConfigHosts '%v' should equal '%v'", actual.IncludeSshConfigHosts, expected.IncludeSshConfigHosts)
	}
	if actual.CheckRemoteForScripts != expected.CheckRemoteForScripts {
		t.Errorf("actual.CheckRemoteForScripts '%v' should equal '%v'", actual.CheckRemoteForScripts, expected.CheckRemoteForScripts)
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
