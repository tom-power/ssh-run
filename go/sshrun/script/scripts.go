package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

type GetScripts = func(host shared.Host, config shared.Config) (string, error)

func GetScriptsFromConf(fs fs.FS) GetScripts {
	return func(host shared.Host, config shared.Config) (string, error) {
		commonScripts, _ := getScriptsFromCommon(fs)
		sharedScripts, _ := getScriptsFromShared(host, fs)
		hostLocalScripts, _ := getScriptsFromHostLocal(host, fs)
		hostRemoteScripts := ""
		if config.CheckRemoteForScripts {
			hostRemoteScripts, _ = getScriptsFromHostRemote(host)
			hostRemoteScripts = " " + hostRemoteScripts
		}
		var out = ""
		out += commonScripts
		out += " " + sharedScripts
		out += " " + hostLocalScripts
		out += hostRemoteScripts
		return removeSh(removeCommandTypes(out)), nil
	}
}

func removeCommandTypes(scripts string) string {
	scripts = strings.ReplaceAll(scripts, ".local", "")
	scripts = strings.ReplaceAll(scripts, ".pty", "")
	scripts = strings.ReplaceAll(scripts, ".x11", "")
	scripts = strings.ReplaceAll(scripts, ".ssh", "")
	return scripts
}

func removeSh(scripts string) string {
	return strings.ReplaceAll(scripts, ".sh", "")
}

func filesToFileNames(files []fs.DirEntry) string {
	filesToFileName := func(dir fs.DirEntry) string { return dir.Name() }
	return strings.Join(shared.Map(files, filesToFileName), " ")
}
