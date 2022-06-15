package config

import (
	"github.com/kevinburke/ssh_config"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

type GetHosts = func() ([]shared.Host, error)

func GetHostsFromSshConfig(path string, fs fs.FS) ([]shared.Host, error) {
	file, err := fs.Open(path)
	if err != nil {
		return []shared.Host{}, err
	}
	bytes, err := getBytes(file)
	if err != nil {
		return []shared.Host{}, err
	}
	config, err := ssh_config.DecodeBytes(bytes)
	if err != nil {
		return []shared.Host{}, err
	}
	return toHosts(config), nil
}

func toHosts(config *ssh_config.Config) []shared.Host {
	var hosts []shared.Host
	for _, host := range config.Hosts {
		for _, pattern := range host.Patterns {
			key := pattern.String()
			name, _ := config.Get(key, "HostName")
			if name != "" {
				user, _ := config.Get(key, "User")
				port, _ := config.Get(key, "Port")
				hosts = append(hosts, shared.Host{
					Host: key,
					User: user,
					Name: name,
					Port: shared.DefaultString(port, "22"),
				})
			}
		}
	}
	return hosts
}
