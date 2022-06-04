package sshrunscripts

import (
	"fmt"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/script"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"strings"
)

type Run = func(
	hostName string,
	scriptName string,
	args []string,
	localUserName string) (string, error)

func GetRun(
	getHost GetHost,
	getScriptPath script.GetScriptPath,
	getScriptContents script.GetScriptContents,
	getCommand GetCommand,
	getScripts script.GetScripts,
	getHosts GetHosts,
) Run {
	return func(
		hostName string,
		scriptName string,
		args []string,
		localUserName string) (string, error) {
		if hostName == "hosts" {
			return echo(getHosts)
		}
		host, err := getHost(hostName, localUserName)
		if err != nil {
			return "", err
		}
		switch scriptName {
		case "explain":
			return hostToString(host)
		case "scripts":
			return echo(func() (string, error) { return getScripts(host) })
		case "ssh":
			return sshCommand(host, getCommand)
		}
		scriptPath, err := getScriptPath(host, scriptName)
		if err != nil {
			return "", err
		}
		scriptContents, err := getScriptContents(host, scriptPath)
		if err != nil {
			return "", err
		}
		command, err := getCommand(commandTypeFromPath(scriptPath), scriptContents, host, args)
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

func commandTypeFromPath(scriptPath string) string {
	fileName := scriptPath[strings.LastIndex(scriptPath, "/")+1:]
	fileNameParts := strings.Split(fileName, ".")
	if len(fileNameParts) == 3 {
		return fileNameParts[1]
	}
	return ""
}
