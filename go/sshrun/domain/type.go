package domain

import (
	"io/fs"
	"strings"
)

func (config Config) Type(fsys fs.FS, host Host, scriptName string) (string, error) {
	path, err := config.Path(fsys, host, scriptName)
	if err != nil {
		return "", err
	}
	return commandTypeFromPath(path), nil
}

func commandTypeFromPath(scriptPath string) string {
	fileName := scriptPath[strings.LastIndex(scriptPath, "/")+1:]
	fileNameParts := strings.Split(fileName, ".")
	if len(fileNameParts) == 3 {
		return fileNameParts[1]
	}
	return ""
}
