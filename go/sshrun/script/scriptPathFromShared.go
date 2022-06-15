package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func scriptPathFromShared(host shared.Host, scriptName string, fsys fs.FS) (string, error) {
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			fromShared, _ := firstFileInSharedDir(hostFile.Name(), scriptName, fsys)
			if fileExists(fsys)(fromShared) {
				return fromShared, nil
			}

		}
	}
	return "", errors.New("nothing in shared")
}

func firstFileInSharedDir(sharedDir string, scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(scriptsPath+"shared/"+sharedDir+"/", scriptName, fs)
}
