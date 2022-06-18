package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"strings"
)

func (fsys FileSys) Type(host domain.Host, scriptName string) (string, error) {
	path, err := fsys.Path(host, scriptName)
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
