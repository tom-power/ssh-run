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
		scripts, err := r.Config.Scripts(r.Fsys, host)
		return echo(scripts), err
	case "ssh", "":
		return host.Ssh(), nil
	}
	scriptContents, err := r.Config.Contents(r.Fsys, host, scriptName)
	if err != nil {
		return "", err
	}
	scriptType, err := r.Config.Type(r.Fsys, host, scriptName)
	if err != nil {
		return "", err
	}
	command, err := host.Command(scriptType, scriptContents, args)
	if err != nil {
		return "", err
	}
	return command, nil
}

func echo(command string) string {
	return "echo " + command
}
