package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (host Host) Command(commandType string, contents string, args []string) (string, error) {
	commandType = cleanup(commandType)
	switch commandType {
	case "ssh":
		return host.Ssh(), nil
	case "local":
		return cleanup(replace(contents, host, args)), nil
	case "":
		return sshRun(host.Ssh(), "", cleanup(replace(contents, host, args))), nil
	case "pty":
		return sshRun(host.Ssh(), "-t", cleanup(replace(contents, host, args))), nil
	case "x11":
		return sshRun(host.Ssh(), "-Y", cleanup(replace(contents, host, args))), nil
	default:
		return "", errors.New("unknown commandType " + commandType)
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
