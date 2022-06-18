package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"github.com/tom-power/ssh-run/sshrun/script"
)

type Runner struct {
	GetCommand GetCommand
	Config     domain.Config
	Script     script.Script
}

func (r Runner) Run(hostName string, scriptName string, args []string) (string, error) {
	if hostName == "hosts" {
		return echo(r.Config.HostNames)
	}
	host, err := r.Config.Host(hostName)
	if err != nil {
		return "", err
	}
	switch scriptName {
	case "explain":
		return host.ToString(), nil
	case "scripts":
		return echo(func() (string, error) { return r.Script.List(host) })
	case "ssh", "":
		return sshCommand(host, r.GetCommand)
	}
	scriptContents, err := r.Script.Contents(host, scriptName)
	if err != nil {
		return "", err
	}
	scriptType, err := r.Script.Type(host, scriptName)
	if err != nil {
		return "", err
	}
	command, err := r.GetCommand(scriptType, scriptContents, host, args)
	if err != nil {
		return "", err
	}
	return command, nil
}

func echo(fn func() (string, error)) (string, error) {
	command, err := fn()
	if err != nil {
		return "", err
	}
	return "echo " + command, nil
}

func sshCommand(host domain.Host, getCommand GetCommand) (string, error) {
	return getCommand("ssh", "", host, []string{""})
}
