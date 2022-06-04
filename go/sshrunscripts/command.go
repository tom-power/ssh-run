package sshrunscripts

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
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
		return sshRun(sshTo(host), "-f", cleanup(withSubs(command, host, args))), nil
	case "sudo":
		return sshRun(sshTo(host), "-t", cleanup(withSubs(command, host, args))), nil
	case "x11":
		return sshRun(sshTo(host), "-Y", cleanup(withSubs(command, host, args))), nil
	default:
		return "", errors.New("unknown commandType " + commandType)
	}
}

func sshTo(host shared.Host) string {
	return fmt.Sprintf("ssh -p %s %s@%s", host.PortSsh, host.User, host.Ip)
}

func sshRun(ssh string, option string, command string) string {
	return fmt.Sprintf("%s %s \"%s\"", ssh, option, command)
}

func withSubs(command string, host shared.Host, args []string) string {
	command = strings.Replace(command, "$ip", host.Ip, -1)
	command = strings.Replace(command, "$host", host.Name, -1)
	command = strings.Replace(command, "$user", host.User, -1)
	command = strings.Replace(command, "$portSsh", host.PortSsh, -1)
	command = strings.Replace(command, "$portTunnel", host.PortTunnel, -1)
	for i := range args {
		command = strings.Replace(command, "$"+strconv.Itoa(i+1), args[i], -1)
	}
	return command
}

func cleanup(command string) string {
	command = strings.Replace(command, "\n", "", -1)
	command = strings.Replace(command, "\\", "", -1)
	command = strings.Trim(command, "\"")
	command = strings.TrimSpace(command)
	space := regexp.MustCompile(`\s+`)
	command = space.ReplaceAllString(command, " ")
	return command
}
