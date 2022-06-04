package script

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"io/ioutil"
)

func scriptPathFromHostLocal(host shared.Host, scriptName string) string {
	script := ""
	hostDir := hostDir(host.Name, homeDir())
	hostFiles, _ := ioutil.ReadDir(hostDir)
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			sharedScript := scriptPathFromShared(hostFile.Name(), scriptName)
			if fileExists(sharedScript) {
				script = sharedScript
			}
			hostSubDirScript := scriptPathFromHostSubDir(host.Name, hostFile.Name(), scriptName)
			if fileExists(hostSubDirScript) {
				script = hostSubDirScript
			}
		}
	}
	hostScript := scriptPathFromHost(scriptsDir(homeDir()), host.Name, scriptName)
	if fileExists(hostScript) {
		script = hostScript
	}
	return script
}

func scriptPathFromHostSubDir(hostsName string, dirName string, scriptName string) string {
	return filePathFromName(scriptsDir(homeDir())+"host/"+hostsName+"/"+dirName+"/", scriptName)
}

func scriptPathFromShared(sharedDir string, scriptName string) string {
	return filePathFromName(scriptsDir(homeDir())+"shared/"+sharedDir+"/", scriptName)
}
