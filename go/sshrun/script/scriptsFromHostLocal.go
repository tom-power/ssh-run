package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func getScriptsFromHostLocal(host shared.Host, fsys fs.FS) (string, error) {
	var files []fs.DirEntry
	err := fs.WalkDir(fsys, hostDir(host.Name), appendFiles(&files))
	if err != nil {
		return "", err
	}
	return filesToFileNames(shared.Filter(files, noKeep)), nil
}

var noKeep = func(file fs.DirEntry) bool { return file.Name() != ".keep" }
