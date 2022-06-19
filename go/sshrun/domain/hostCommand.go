package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func (h Host) Command(scriptType ScriptType, contents string, args []string) (string, error) {
	command := h.command(contents, args)
	switch scriptType {
	case Default:
		return h.SshWith(command, ""), nil
	case Pty:
		return h.SshWith(command, "-t"), nil
	case X11:
		return h.SshWith(command, "-Y"), nil
	case Local:
		return command, nil
	default:
		return "", errors.New("unknown scriptType " + scriptType.String())
	}
}

func (h Host) command(contents string, args []string) string {
	return cleanup(h.replace(contents, args))
}

func (h Host) replace(command string, args []string) string {
	command = strings.Replace(command, "$hostName", h.Name, -1)
	command = strings.Replace(command, "$h", h.Host, -1)
	command = strings.Replace(command, "$user", h.User, -1)
	command = strings.Replace(command, "$portTunnel", h.PortTunnel, -1)
	command = strings.Replace(command, "$port", h.Port, -1)
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
