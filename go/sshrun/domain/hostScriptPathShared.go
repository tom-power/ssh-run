package domain

import (
	"errors"
	"io/fs"
)

func (host Host) pathShared(fsys fs.FS, scriptName string) (string, error) {
	hostFiles, err := host.Files(fsys)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			scriptPathShared, _ := findScriptPathInShared(fsys, scriptName, hostFile)
			if fileExists(fsys)(scriptPathShared) {
				return scriptPathShared, nil
			}

		}
	}
	return "", errors.New("nothing in shared")
}

func findScriptPathInShared(fsys fs.FS, scriptName string, hostFile fs.DirEntry) (string, error) {
	return firstPathInDir(fsys, scriptsPath+"shared/"+hostFile.Name()+"/", scriptName)
}
