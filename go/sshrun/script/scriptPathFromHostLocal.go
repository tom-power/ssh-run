package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func scriptPathFromHostLocal(host shared.Host, scriptName string, fsys fs.FS) (string, error) {
	script := ""
	hostDir := hostDirRel(host.Name)
	hostFiles, err := fs.ReadDir(fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			fromShared, _ := scriptPathFromShared(hostFile.Name(), scriptName, fsys)
			shared.UpdateIf(&script, fromShared, fileExistsFs(fsys))
			subDir, _ := scriptPathFromHostSubDir(host.Name, hostFile.Name(), scriptName, fsys)
			shared.UpdateIf(&script, subDir, fileExistsFs(fsys))
		}
	}
	fromHost, err := scriptPathFromHost(scriptsPath, host.Name, scriptName, fsys)
	shared.UpdateIf(&script, fromHost, fileExistsFs(fsys))
	return script, nil
}

func scriptPathFromHostSubDir(hostsName string, dirName string, scriptName string, fs fs.FS) (string, error) {
	return filePathFromNameFs(scriptsPath+"host/"+hostsName+"/"+dirName+"/", scriptName, fs)
}

func scriptPathFromShared(sharedDir string, scriptName string, fs fs.FS) (string, error) {
	return filePathFromNameFs(scriptsPath+"shared/"+sharedDir+"/", scriptName, fs)
}
