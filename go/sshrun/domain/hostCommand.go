package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (host Host) Command(scriptType ScriptType, contents string, args []string) (string, error) {
	switch scriptType {
	case Ssh:
		return host.Ssh(), nil
	case Local:
		return cleanup(replace(contents, host, args)), nil
	case Default:
		return sshRun(host.Ssh(), "", cleanup(replace(contents, host, args))), nil
	case Pty:
		return sshRun(host.Ssh(), "-t", cleanup(replace(contents, host, args))), nil
	case X11:
		return sshRun(host.Ssh(), "-Y", cleanup(replace(contents, host, args))), nil
	default:
		return "", errors.New("unknown scriptType " + scriptType.String())
	}
}

func (host Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", host.Port, host.User, host.Host)
}

func sshRun(ssh string, option string, command string) string {
	if option == "" {
		return fmt.Sprintf("%s \"%s\"", ssh, command)
	}
	return fmt.Sprintf("%s %s \"%s\"", ssh, option, command)
}

func replace(command string, host Host, args []string) string {
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

func (host Host) withSubs(command string, args []string) string {
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
