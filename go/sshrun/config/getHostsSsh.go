package config

import (
	"io/fs"

	"github.com/kevinburke/ssh_config"
	fp "github.com/rjNemo/underscore"
	"github.com/tom-power/ssh-run/sshrun/domain"
	"github.com/tom-power/ssh-run/sshrun/fp/pred"
)

func (c ConfigFs) getHostsFromSshConfig() ([]domain.Host, error) {
	file, err := fs.ReadFile(c.Fsys, c.SshPath)
	if err != nil {
		return []domain.Host{}, err
	}
	config, err := ssh_config.DecodeBytes(file)
	if err != nil {
		return []domain.Host{}, err
	}
	return toHosts(config), nil
}

func toHosts(config *ssh_config.Config) []domain.Host {
	var hosts []domain.Host
	for _, host := range config.Hosts {
		for _, pattern := range host.Patterns {
			key := pattern.String()
			name, _ := config.Get(key, "HostName")
			if name != "" {
				user, _ := config.Get(key, "User")
				port, _ := config.Get(key, "Port")
				hosts = append(hosts, domain.Host{
					Ip:   key,
					User: user,
					Name: name,
					Port: fp.Ternary(pred.IsNotEmpty(port), port, "22"),
				})
			}
		}
	}
	return hosts
}
