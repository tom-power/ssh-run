package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (fsys FileSys) Contents(host shared.Host, scriptPath string) (string, error) {
	script, err := fsys.contentsFromHostLocal(scriptPath)
	if fsys.Config.CheckRemoteForScripts {
		script, err = contentsFromHostRemote(host, scriptPath)
	}
	if err != nil {
		return "", err
	}
	return script, nil
}

func (fsys FileSys) contentsFromHostLocal(scriptPath string) (string, error) {
	dat, err := fs.ReadFile(fsys.Fsys, scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func contentsFromHostRemote(host shared.Host, scriptPath string) (string, error) {
	return runCommandOn(host, "cat "+scriptPath)
}
