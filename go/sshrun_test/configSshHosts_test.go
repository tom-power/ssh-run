package sshrun

import (
	"testing"
	"testing/fstest"

	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/domain"
)

func Test_fileSysConfigIncludeSshConfigHosts(t *testing.T) {

	t.Run("get hosts from ssh config when IncludeSshConfigHosts fs", func(t *testing.T) {
		var testFsConfig = fstest.MapFS{
			"config/ssh-run.config.yml": {Data: []byte(configYamlHostsFromSshText)},
			".ssh/config":               {Data: []byte(hostsSshConfigText)},
		}

		sys := config.ConfigFs{
			Fsys:      testFsConfig,
			ConfigDir: "config",
			SshPath:   ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{
			Hosts:                 append(expectedConfigFromYaml.Hosts, sshConfigHosts...),
			IncludeSshConfigHosts: true,
		})
	})
	t.Run("don't get hosts from ssh config when not IncludeSshConfigHosts fs", func(t *testing.T) {
		var testFsConfig = fstest.MapFS{
			"config/ssh-run.config.yml": {Data: []byte(configYamlNoHostsFromSshText)},
			".ssh/config":               {Data: []byte(hostsSshConfigText)},
		}

		sys := config.ConfigFs{
			Fsys:      testFsConfig,
			ConfigDir: "config",
			SshPath:   ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{
			Hosts:                 expectedConfigFromYaml.Hosts,
			IncludeSshConfigHosts: false,
		})
	})
	t.Run("get hosts from ssh config when no hosts in yaml", func(t *testing.T) {
		var testFsConfig = fstest.MapFS{
			"config.yml":  {Data: []byte(``)},
			".ssh/config": {Data: []byte(hostsSshConfigText)},
		}

		sys := config.ConfigFs{
			Fsys:      testFsConfig,
			ConfigDir: "config",
			SshPath:   ".ssh/config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, domain.Config{Hosts: sshConfigHosts})
	})
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

var configYamlHostsFromSshText = `
includeSshConfigHosts: true
hosts:
  - ip: 192.0.2.1
    user: testUser1
    name: testName1
    port: 22
  - ip: 192.0.2.2
    user: testUser2
    name: testName2
    port: 23
    portTunnel: 1081
  - ip: 192.0.2.3
    user: testUser3
    name: testName3
    port: 24`

var configYamlNoHostsFromSshText = `
includeSshConfigHosts: false
hosts:
  - ip: 192.0.2.1
    user: testUser1
    name: testName1
    port: 22
  - ip: 192.0.2.2
    user: testUser2
    name: testName2
    port: 23
    portTunnel: 1081
  - ip: 192.0.2.3
    user: testUser3
    name: testName3
    port: 24`

var sshConfigHosts = []domain.Host{{
	Name: "localhost",
	User: "test",
	Ip:   "localhost",
	Port: "22",
}, {
	Name: "github.com",
	User: "git",
	Ip:   "github.com",
	Port: "22",
}}
