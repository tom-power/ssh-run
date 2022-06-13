package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/ioutil"
)

func scriptPathFromHostLocal(host shared.Host, scriptName string) string {
	script := ""
	hostDir := hostDir(host.Name, homeDir())
	hostFiles, _ := ioutil.ReadDir(hostDir)
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			shared.UpdateIf(&script, scriptPathFromShared(hostFile.Name(), scriptName), fileExists)
			shared.UpdateIf(&script, scriptPathFromHostSubDir(host.Name, hostFile.Name(), scriptName), fileExists)
		}
	}
	shared.UpdateIf(&script, scriptPathFromHost(scriptsDir(homeDir()), host.Name, scriptName), fileExists)
	return script
}

func scriptPathFromHostSubDir(hostsName string, dirName string, scriptName string) string {
	return filePathFromName(scriptsDir(homeDir())+"host/"+hostsName+"/"+dirName+"/", scriptName)
}

func scriptPathFromShared(sharedDir string, scriptName string) string {
	return filePathFromName(scriptsDir(homeDir())+"shared/"+sharedDir+"/", scriptName)
}
