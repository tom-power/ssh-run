package script

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"strings"
)

func getScriptsFromHost(host shared.Host) (string, error) {
	local, err := scriptsFromHostLocal(host)
	if err != nil {
		return "", err
	}
	var remote = ""
	if host.CheckForScripts {
		remote, err = scriptsFromHostRemote(host)
		if err == nil {
			remote = " " + remote
		}
	}
	return toString(shared.Unique(toSlice(local + remote))), nil
}

func toSlice(str string) []string {
	return strings.Split(str, " ")
}

func toString(strSlice []string) string {
	return strings.Join(strSlice, " ")
}
