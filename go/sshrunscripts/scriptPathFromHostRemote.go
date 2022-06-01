package sshrunscripts

import (
	"bytes"
)

func scriptFromHostRemote(host Host, scriptName string) (string, error) {
	session, err := getSession(host)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	session.Stdout = &buff
	scriptPath := scriptPathFromHost(scriptsDir("/home/"+host.User), "localhost", scriptName)
	if err := session.Run("cat " + scriptPath); err != nil {
		return "", err
	}
	return buff.String(), nil
}

