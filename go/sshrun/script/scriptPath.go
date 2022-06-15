package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"os/user"
)

type GetScriptPath = func(host shared.Host, scriptName string, config shared.Config) (string, error)

func GetScriptPathFromConf(fs fs.FS) GetScriptPath {
	return func(host shared.Host, scriptName string, config shared.Config) (string, error) {
		script := ""
		fileExistsFs := fileExistsFs(fs)
		common, _ := scriptPathFromCommon(scriptName, fs)
		shared.UpdateIf(&script, common, fileExistsFs)
		fromHost, _ := getScriptPathFromHost(host, scriptName, config, fs)
		shared.UpdateIf(&script, fromHost, fileExistsFs)
		if script == "" {
			return "", errors.New("couldn't find script \"" + scriptName + ".sh\" for host \"" + host.Name + "\"")
		}
		return script, nil
	}
}

const scriptsPath = ".config/ssh-run/scripts/"

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func hostDir(hostsName string) string {
	return scriptsPath + "host/" + hostsName
}

func scriptPathFromHost(scriptsDir string, hostsName string, scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(scriptsDir+"host/"+hostsName+"/", scriptName, fs)
}

func commonDir() string {
	return scriptsPath + "common"
}

func scriptPathFromCommon(scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(commonDir()+"/", scriptName, fs)
}

func firstFileInDir(dir string, name string, fsys fs.FS) (string, error) {
	matches, err := fs.Glob(fsys, dir+name+".*")
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", errors.New("no match")
	}
	return matches[0], nil
}
