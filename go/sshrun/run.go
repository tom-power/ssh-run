package sshrun

import (
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"strings"
)

type Run = func(
	hostName string,
	scriptName string,
	args []string) (string, error)

func GetRun(
	getCommand GetCommand,
	config shared.Config,
	script script.Script,
) Run {
	return func(
		hostName string,
		scriptName string,
		args []string) (string, error) {
		if hostName == "hosts" {
			return echo(config.HostNames)
		}
		host, err := config.Host(hostName)
		if err != nil {
			return "", err
		}
		switch scriptName {
		case "explain":
			return hostToString(host)
		case "scripts":
			return echo(func() (string, error) { return script.List(host) })
		case "ssh", "":
			return sshCommand(host, getCommand)
		}
		scriptContents, err := script.Contents(host, scriptName)
		if err != nil {
			return "", err
		}
		scriptType, err := script.Type(host, scriptName)
		if err != nil {
			return "", err
		}
		command, err := getCommand(scriptType, scriptContents, host, args)
		if err != nil {
			return "", err
		}
		return command, nil
	}
}

func hostToString(host shared.Host) (string, error) {
	return strings.ReplaceAll(fmt.Sprintf("%#v", host), "shared.", ""), nil
}

func echo(fn func() (string, error)) (string, error) {
	command, err := fn()
	if err != nil {
		return "", err
	}
	return "echo " + command, nil
}

func sshCommand(host shared.Host, getCommand GetCommand) (string, error) {
	command, err := getCommand("ssh", "", host, []string{""})
	if err != nil {
		return "", err
	}
	return command, nil
}
