package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"io/fs"
)

type Runner struct {
	Config domain.Config
	Fsys   fs.FS
}

func (r Runner) Run(hostName string, scriptName string, args []string) (string, error) {
	if hostName == "hosts" {
		names, err := r.Config.HostNames()
		return echo(names), err
	}
	host, err := r.Config.Host(hostName)
	if err != nil {
		return "", err
	}
	switch scriptName {
	case "explain":
		return host.ToString(), nil
	case "scripts":
		scripts, err := host.Scripts(r.Fsys)
		return echo(scripts), err
	case "ssh", "":
		return host.Ssh(), nil
	}
	script, err := host.Script(r.Fsys, scriptName)
	if err != nil {
		return "", err
	}
	command, err := host.Command(script, args)
	if err != nil {
		return "", err
	}
	return command, nil
}

func echo(command string) string {
	return "echo " + command
}
