package script

import (
	"errors"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"os/user"
	"path/filepath"
)

type GetScriptPath = func(host shared.Host, scriptName string) (string, error)

var GetScriptPathFromConf = func(host shared.Host, scriptName string) (string, error) {
	script := ""
	shared.UpdateIf(&script, scriptPathFromCommon(scriptName), fileExists)
	shared.UpdateIf(&script, getScriptPathFromHost(host, scriptName), fileExists)
	if script == "" {
		return "", errors.New("couldn't find script \"" + scriptName + ".sh\" for host \"" + host.Name + "\"")
	}
	return script, nil
}

const scriptsPath = "/.config/ssh-run-scripts/scripts/"

func scriptsDir(homeDir string) string {
	return homeDir + scriptsPath
}

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func hostDir(hostsName string, homeDir string) string {
	return scriptsDir(homeDir) + "host/" + hostsName + "/"
}

func scriptPathFromHost(scriptsDir string, hostsName string, scriptName string) string {
	return filePathFromName(scriptsDir+"host/"+hostsName+"/", scriptName)
}

func commonDir() string {
	return scriptsDir(homeDir()) + "common/"
}

func scriptPathFromCommon(scriptName string) string {
	return filePathFromName(commonDir(), scriptName)
}

func filePathFromName(dir string, name string) string {
	matches, err := filepath.Glob(dir + name + ".*")
	if err != nil || len(matches) == 0 {
		return ""
	}
	return matches[0]
}
