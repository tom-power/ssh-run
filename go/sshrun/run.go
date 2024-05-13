package sshrun

import (
	_ "embed"
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/domain"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type Runner struct {
	Config domain.Config
	Fsys   fs.FS
}

//go:embed embed/help.txt
var help string

func (r Runner) Run(hostName string, scriptName string, args []string) (string, error) {
	if hostName == "" || shared.Intersect([]string{"--help", "-h"}, args) {
		return echo(help), nil
	}

	switch hostName {
	case "hosts":
		names, err := r.Config.HostNames()
		return echo("localhost " + names), err
	}

	host, err := r.getHost(hostName)
	if err != nil {
		return "", err
	}

	switch scriptName {
	case "scripts":
		scripts, err := host.Scripts(r.Fsys)
		return echo(scripts), err
	case "ssh":
		return host.Ssh(), nil
	case "":
		if shared.Any(args, "--explain") {
			return host.ToString(), nil
		}
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

func (r Runner) getHost(hostName string) (domain.Host, error) {
	if hostName == "localhost" && r.Config.LocalhostIs != "" {
		hostName = r.Config.LocalhostIs
	}
	host, err := r.Config.Host(hostName)
	if err != nil {
		return host, err
	}
	return host, nil
}

func echo(command string) string {
	return "echo " + command
}
