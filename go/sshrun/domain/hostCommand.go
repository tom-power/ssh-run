package domain

import (
	"errors"
	"regexp"
	"strings"
)

func (h Host) Command(script Script) (string, error) {
	command := command(script.Contents, h)
	switch script.Type {
	case Default:
		return h.SshWith(command, ""), nil
	case Pty:
		return h.SshWith(command, "-t"), nil
	case X11:
		return h.SshWith(command, "-Y"), nil
	case Local:
		return command, nil
	default:
		return "", errors.New("unknown scriptType " + script.Type.String())
	}
}

func command(contents string, host Host) string {
	return cleanup(replace(contents, host))
}

func replace(command string, host Host) string {
	command = strings.Replace(command, "$hostName", host.Name, -1)
	command = strings.Replace(command, "$ip", host.Ip, -1)
	command = strings.Replace(command, "$user", host.User, -1)
	command = strings.Replace(command, "$portTunnel", host.PortTunnel, -1)
	command = strings.Replace(command, "$port", host.Port, -1)
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
