package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func scriptPathFromHostLocal(host shared.Host, scriptName string, fsys fs.FS) (string, error) {
	script := ""
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			subDir, _ := firstFileInHostSubDir(host.Name, hostFile.Name(), scriptName, fsys)
			shared.UpdateIf(&script, subDir, fileExists(fsys))
		}
	}
	fromHost, err := scriptPathFromHost(scriptsPath, host.Name, scriptName, fsys)
	shared.UpdateIf(&script, fromHost, fileExists(fsys))
	return script, nil
}

func firstFileInHostSubDir(hostsName string, dirName string, scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(scriptsPath+"host/"+hostsName+"/"+dirName+"/", scriptName, fs)
}

func scriptPathFromHost(scriptsDir string, hostsName string, scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(scriptsDir+"host/"+hostsName+"/", scriptName, fs)
}
