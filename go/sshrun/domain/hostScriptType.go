package domain

import (
	"io/fs"
	"strings"
)

func (h Host) Type(fsys fs.FS, scriptName string) (ScriptType, error) {
	path, err := h.Path(fsys, scriptName)
	if err != nil {
		return Default, err
	}
	return Default.Parse(scriptTypeFrom(path)), nil
}

func scriptTypeFrom(scriptPath string) string {
	fileName := scriptPath[strings.LastIndex(scriptPath, "/")+1:]
	fileNameParts := strings.Split(fileName, ".")
	if len(fileNameParts) == 3 {
		return fileNameParts[1]
	}
	return ""
}
