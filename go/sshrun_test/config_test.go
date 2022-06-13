package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_config(t *testing.T) {
	sshConfigHosts := []shared.Host{{
		Name: "testNameExtra",
		User: "testUserExtra",
		Host: "192.0.2.9",
		Port: "22",
	}}
	getHostsFromSshConfig := func() ([]shared.Host, error) {
		return sshConfigHosts, nil
	}

	t.Run("get hosts from ssh config when IncludeSshConfig", func(t *testing.T) {
		configFromYaml := shared.Config{
			Hosts:            expectedConfigFromYaml.Hosts,
			IncludeSshConfig: true,
		}
		getConfigFromYaml := func() (shared.Config, error) {
			return configFromYaml, nil
		}

		actual, _ := config.GetConfigFromFileSystem(getHostsFromSshConfig, getConfigFromYaml)

		assertConfigEqual(t, actual, shared.Config{
			Hosts:            append(configFromYaml.Hosts, sshConfigHosts...),
			IncludeSshConfig: configFromYaml.IncludeSshConfig,
		})
	})
	t.Run("don't get hosts from ssh config when not IncludeSshConfig", func(t *testing.T) {
		configFromYaml := shared.Config{
			Hosts:            expectedConfigFromYaml.Hosts,
			IncludeSshConfig: false,
		}
		getConfigFromYaml := func() (shared.Config, error) {
			return configFromYaml, nil
		}

		actual, _ := config.GetConfigFromFileSystem(getHostsFromSshConfig, getConfigFromYaml)

		assertConfigEqual(t, actual, shared.Config{
			Hosts:            append(configFromYaml.Hosts),
			IncludeSshConfig: configFromYaml.IncludeSshConfig,
		})
	})
	t.Run("get hosts from ssh config when nothing in yaml", func(t *testing.T) {
		getConfigFromYaml := func() (shared.Config, error) {
			return shared.Config{}, nil
		}

		actual, _ := config.GetConfigFromFileSystem(getHostsFromSshConfig, getConfigFromYaml)

		assertConfigEqual(t, actual, shared.Config{Hosts: sshConfigHosts})
	})
}
