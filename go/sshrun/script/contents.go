package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"io/fs"
)

func (fsys FileSys) Contents(host domain.Host, scriptName string) (string, error) {
	path, err := fsys.Path(host, scriptName)
	if err != nil {
		return "", err
	}
	script, err := fsys.contentsFromHostLocal(path)
	if fsys.Config.CheckRemoteForScripts {
		script, err = contentsFromHostRemote(host, path)
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

func contentsFromHostRemote(host domain.Host, scriptPath string) (string, error) {
	return runCommandOn(host, "cat "+scriptPath)
}
