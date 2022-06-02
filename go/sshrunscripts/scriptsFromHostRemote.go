package sshrunscripts

import (
	"bytes"
)

func scriptsFromHostLocalRemote(host Host) (string, error) {
	session, err := getSession(host)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	session.Stdout = &buff
	scriptsDir := scriptsDir("/home/" + host.User + "host/localhost")
	if err := session.Run("find " + scriptsDir + " -type f -printf '%%f ' | sed \"s/.sh//g\""); err != nil {
		return "", err
	}
	return buff.String(), nil
}
