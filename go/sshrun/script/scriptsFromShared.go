package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func getScriptsFromShared(host shared.Host, fsys fs.FS) (string, error) {
	var files []fs.DirEntry
	hostFiles, _ := fs.ReadDir(fsys, hostDir(host.Name))
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			sharedDir := scriptsPath + "shared/" + hostFile.Name()
			if fileExists(fsys)(sharedDir) {
				err := fs.WalkDir(fsys, sharedDir, appendFiles(&files))
				if err != nil {
					return "", err
				}
			}
		}
	}
	return filesToFileNames(files), nil
}

func appendFiles(files *[]fs.DirEntry) func(string, fs.DirEntry, error) error {
	return func(path string, info fs.DirEntry, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
