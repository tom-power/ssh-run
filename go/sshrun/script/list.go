package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

func (fsys FileSys) List(host shared.Host) (string, error) {
	commonScripts, _ := fsys.listCommon()
	sharedScripts, _ := fsys.listShared(host)
	hostLocalScripts, _ := fsys.listHostLocal(host)
	hostRemoteScripts := ""
	if fsys.Config.CheckRemoteForScripts {
		hostRemoteScripts, _ = listHostRemote(host)
		hostRemoteScripts = " " + hostRemoteScripts
	}
	var out = ""
	out += commonScripts
	out += " " + sharedScripts
	out += " " + hostLocalScripts
	out += hostRemoteScripts
	return removeSh(removeCommandTypes(out)), nil
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
