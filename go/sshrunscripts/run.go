package sshrunscripts

import (
	"strings"
)

func Run(
	getHost GetHost,
	getScriptPath GetScriptPath,
	getScript GetScript,
	getCommand GetCommand,
	getScripts GetScripts,
	getHosts GetHosts,
) func(string, string, []string, string) (string, error) {
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
		if scriptName == "scripts" {
			return echo(func() (string, error) { return getScripts(host) })
		}
		if scriptName == "ssh" {
			return sshCommand(host, getCommand)
		}
		var prefix = ""
		if scriptName == "explain" {
			scriptName = args[0]
			prefix = "echo "
		}
		scriptPath, err := getScriptPath(host, scriptName)
		if err != nil {
			return "", err
		}
		scriptContents, err := getScript(scriptPath)
		if err != nil {
			return "", err
		}
		command, err := getCommand(commandTypeFromPath(scriptPath), scriptContents, host, args)
		if err != nil {
			return "", err
		}
		return prefix + command, nil
	}
}

func echo(fn func() (string, error)) (string, error) {
	command, err := fn()
	if err != nil {
		return "", err
	}
	return "echo " + command, nil
}

func sshCommand(host Host, getCommand GetCommand) (string, error) {
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
