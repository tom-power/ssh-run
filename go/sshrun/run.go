package sshrun

import (
	_ "embed"
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/domain"
)

type Runner struct {
	Config domain.Config
	Fsys   fs.FS
}

//go:embed embed/help.txt
var helpTxt string

func (r Runner) Run(hostName string, scriptName string, flags RunFlags) (string, error) {
	switch {
	case hostName == "" && scriptName == "":
		if flags.Hosts {
			names, err := r.Config.HostNames()
			return "localhost " + names, err
		}
	case hostName != "":
		host, err := r.getHost(hostName)
		if err != nil {
			return "", err
		}
		switch scriptName {
		case "":
			if flags.Scripts {
				return host.Scripts(r.Fsys)
			}
			if flags.Explain {
				return host.ToString(), nil
			}
			if flags.Ssh {
				return host.Ssh(), nil
			}
		default:
			script, err := host.Script(r.Fsys, scriptName)
			if err != nil {
				return "", err
			}
			command, err := host.Command(script)
			if err != nil {
				return "", err
			}
			return command, nil
		}
	}
	return helpTxt, nil
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
	Help    bool
	Hosts   bool
	Scripts bool
	Explain bool
	Ssh     bool
}
