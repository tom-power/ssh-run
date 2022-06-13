package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_config(t *testing.T) {
	t.Run("can get config from files system with yaml config", func(t *testing.T) {
		sshConfigHosts := []shared.Host{{
			Name: "testNameExtra",
			User: "testUserExtra",
			Host: "192.0.2.9",
			Port: "22",
		}}
		getHostsFromSshConfig := func() ([]shared.Host, error) {
			return sshConfigHosts, nil
		}
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
}
