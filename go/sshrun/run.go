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
var helpTxt string

func (r Runner) Run(hostName string, scriptName string, flags RunFlags) (string, error) {
	if (hostName == "" && flags.empty()) || flags.Help {
		return helpTxt, nil
	}

	if flags.Hosts {
		names, err := r.Config.HostNames()
		return "localhost " + names, err
	}

	host, err := r.getHost(hostName)
	if err != nil {
		return "", err
	}

	switch scriptName {
	case "ssh":
		return host.Ssh(), nil
	case "":
		if flags.Explain {
			return host.ToString(), nil
		}
		if flags.Scripts {
			scripts, err := host.Scripts(r.Fsys)
			return scripts, err
		}
	}
	script, err := host.Script(r.Fsys, scriptName)
	if err != nil {
		return "", err
	}

	command, err := host.Command(script, flags.ScriptArgs)
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

type RunFlags struct {
	Help       bool
	Hosts      bool
	Scripts    bool
	Explain    bool
	ScriptArgs []string
}

func (r *RunFlags) empty() bool {
	flags := []bool{
		r.Explain, r.Help, r.Hosts, r.Scripts, (len(r.ScriptArgs) > 0),
	}
	return shared.All(flags, false)
}
