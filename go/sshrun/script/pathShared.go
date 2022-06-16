package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (fsys FileSys) pathShared(host shared.Host, scriptName string) (string, error) {
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys.Fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			fromShared, _ := fsys.firstFileInSharedDir(hostFile.Name(), scriptName)
			if fsys.fileExists()(fromShared) {
				return fromShared, nil
			}

		}
	}
	return "", errors.New("nothing in shared")
}

func (fsys FileSys) firstFileInSharedDir(sharedDir string, scriptName string) (string, error) {
	return fsys.firstFileInDir(scriptsPath+"shared/"+sharedDir+"/", scriptName)
}
