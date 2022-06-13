package sshrun

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"regexp"
	"strconv"
	"strings"
)

type GetCommand = func(commandType string, command string, host shared.Host, args []string) (string, error)

var GetCommandSsh = func(
	commandType string,
	command string,
	host shared.Host,
	args []string) (string, error) {
	commandType = cleanup(commandType)
	switch commandType {
	case "ssh":
		return sshTo(host), nil
	case "local":
		return cleanup(withSubs(command, host, args)), nil
	case "":
		return sshRun(sshTo(host), "", cleanup(withSubs(command, host, args))), nil
	case "pty":
		return sshRun(sshTo(host), "-t", cleanup(withSubs(command, host, args))), nil
	case "x11":
		return sshRun(sshTo(host), "-Y", cleanup(withSubs(command, host, args))), nil
	default:
		return "", errors.New("unknown commandType " + commandType)
	}
}

func sshTo(host shared.Host) string {
	return fmt.Sprintf("ssh -p %s %s@%s", host.Port, host.User, host.Host)
}

func sshRun(ssh string, option string, command string) string {
	if option == "" {
		return fmt.Sprintf("%s \"%s\"", ssh, command)
	}
	return fmt.Sprintf("%s %s \"%s\"", ssh, option, command)
}

func withSubs(command string, host shared.Host, args []string) string {
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
