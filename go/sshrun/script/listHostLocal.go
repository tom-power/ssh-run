package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (fsys FileSys) listHostLocal(host shared.Host) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	err := fs.WalkDir(fsys.Fsys, hostDir(host.Name), appendFiles(&files.Files))
	if err != nil {
		return "", err
	}
	return files.names(), nil
}
