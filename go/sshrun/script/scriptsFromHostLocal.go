package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"os"
	"path/filepath"
)

func scriptsFromHostLocal(host shared.Host) (string, error) {
	var files []os.FileInfo
	err := filepath.Walk(hostDir(host.Name, homeDir()), appendFiles(&files))
	return filesToFileNames(shared.Filter(files, noKeep)), err
}

var noKeep = func(file os.FileInfo) bool { return file.Name() != ".keep" }
