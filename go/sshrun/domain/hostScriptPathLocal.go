package domain

import (
	"io/fs"
)

func (h Host) pathLocal(fsys fs.FS, scriptName string) (string, error) {
	script := ""
	hostFiles, err := h.Files(fsys)
	if err != nil {
		return "", err
	}
	for _, file := range hostFiles {
		if file.IsDir() {
			scriptPathSubDir, _ := pathInDir(fsys, h.Dir()+"/"+file.Name()+"/", scriptName)
			replaceIf(&script, scriptPathSubDir, fileExists(fsys))
		}
	}
	scriptPath, _ := pathInDir(fsys, h.Dir()+"/", scriptName)
	replaceIf(&script, scriptPath, fileExists(fsys))
	return script, nil
}

func replaceIf[V any](value *V, newValue V, predicate func(V) bool) {
	if predicate(newValue) {
		*value = newValue
	}
}
