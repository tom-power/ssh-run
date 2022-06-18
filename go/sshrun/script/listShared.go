package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"io/fs"
)

func (fsys FileSys) listShared(host domain.Host) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	hostFiles, _ := fs.ReadDir(fsys.Fsys, hostDir(host.Name))
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			sharedDir := scriptsPath + "shared/" + hostFile.Name()
			if fsys.fileExists()(sharedDir) {
				err := fs.WalkDir(fsys.Fsys, sharedDir, appendFiles(&files.Files))
				if err != nil {
					return "", err
				}
			}
		}
	}
	return files.names(), nil
}

func appendFiles(files *[]fs.DirEntry) func(string, fs.DirEntry, error) error {
	return func(path string, info fs.DirEntry, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
