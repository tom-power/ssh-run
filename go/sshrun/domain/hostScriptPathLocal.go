package domain

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (host Host) pathLocal(fsys fs.FS, scriptName string) (string, error) {
	script := ""
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			subDir, _ := firstFileInHostSubDir(fsys, host.Name, hostFile.Name(), scriptName)
			shared.ReplaceIf(&script, subDir, fileExists(fsys))
		}
	}
	fromHost, err := scriptPathFromHost(fsys, scriptsPath, host.Name, scriptName)
	shared.ReplaceIf(&script, fromHost, fileExists(fsys))
	return script, nil
}

func firstFileInHostSubDir(fsys fs.FS, hostsName string, dirName string, scriptName string) (string, error) {
	return firstFileInDir(fsys, scriptsPath+"host/"+hostsName+"/"+dirName+"/", scriptName)
}

func scriptPathFromHost(fsys fs.FS, scriptsDir string, hostsName string, scriptName string) (string, error) {
	return firstFileInDir(fsys, scriptsDir+"host/"+hostsName+"/", scriptName)
}
