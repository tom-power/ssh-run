package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_config(t *testing.T) {
	t.Run("can get config from files system with ssh config", func(t *testing.T) {
		hostsFromSshConfigYamlConfig := shared.Config{
			Hosts:              expectedConfigFromYaml.Hosts,
			HostsFromSshConfig: true,
		}
		getFromYaml := func() (shared.Config, error) {
			return hostsFromSshConfigYamlConfig, nil
		}
		extraHost := shared.Host{
			Name: "testNameExtra",
			User: "testUserExtra",
			Host: "192.0.2.9",
			Port: "22",
		}
		getFromSshConfig := func() ([]shared.Host, error) {
			return []shared.Host{extraHost}, nil
		}

		actual, _ := config.GetConfigFromFileSystem(getFromYaml, getFromSshConfig)

		assertConfigEqual(t, actual, shared.Config{
			Hosts:              append(hostsFromSshConfigYamlConfig.Hosts, extraHost),
			HostsFromSshConfig: hostsFromSshConfigYamlConfig.HostsFromSshConfig,
		})
	})
}
