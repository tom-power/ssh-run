package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func (host Host) Command(scriptType ScriptType, contents string, args []string) (string, error) {
	command := host.command(contents, args)
	switch scriptType {
	case Default:
		return host.SshWith(command, ""), nil
	case Pty:
		return host.SshWith(command, "-t"), nil
	case X11:
		return host.SshWith(command, "-Y"), nil
	case Local:
		return command, nil
	default:
		return "", errors.New("unknown scriptType " + scriptType.String())
	}
}

func (host Host) command(contents string, args []string) string {
	return cleanup(host.replace(contents, args))
}

func (host Host) replace(command string, args []string) string {
	command = strings.Replace(command, "$hostName", host.Name, -1)
	command = strings.Replace(command, "$host", host.Host, -1)
	command = strings.Replace(command, "$user", host.User, -1)
	command = strings.Replace(command, "$portTunnel", host.PortTunnel, -1)
	command = strings.Replace(command, "$port", host.Port, -1)
	for i := range args {
		command = strings.Replace(command, "$"+strconv.Itoa(i+1), args[i], -1)
	}
	return command
}

func cleanup(command string) string {
	command = strings.Replace(command, "\n", " ", -1)
	command = strings.Replace(command, "\\", "", -1)
	command = strings.TrimSpace(command)
	space := regexp.MustCompile(`\s+`)
	command = space.ReplaceAllString(command, " ")
	return command
}
