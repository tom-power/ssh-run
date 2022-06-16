package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (fsys FileSys) pathFromHostLocal(host shared.Host, scriptName string) (string, error) {
	script := ""
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys.Fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			subDir, _ := fsys.firstFileInHostSubDir(host.Name, hostFile.Name(), scriptName)
			shared.UpdateIf(&script, subDir, fsys.fileExists())
		}
	}
	fromHost, err := fsys.scriptPathFromHost(scriptsPath, host.Name, scriptName)
	shared.UpdateIf(&script, fromHost, fsys.fileExists())
	return script, nil
}

func (fsys FileSys) firstFileInHostSubDir(hostsName string, dirName string, scriptName string) (string, error) {
	return fsys.firstFileInDir(scriptsPath+"host/"+hostsName+"/"+dirName+"/", scriptName)
}

func (fsys FileSys) scriptPathFromHost(scriptsDir string, hostsName string, scriptName string) (string, error) {
	return fsys.firstFileInDir(scriptsDir+"host/"+hostsName+"/", scriptName)
}
