package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

func getScriptsFromHost(host shared.Host, config shared.Config, fs fs.FS) (string, error) {
	local, err := scriptsFromHostLocal(host, fs)
	if err != nil {
		return "", err
	}
	var remote = ""
	if config.CheckRemoteForScripts {
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
