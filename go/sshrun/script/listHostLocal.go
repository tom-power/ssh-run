package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"io/fs"
)

func (fsys FileSys) listHostLocal(host domain.Host) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	err := fs.WalkDir(fsys.Fsys, hostDir(host.Name), appendFiles(&files.Files))
	if err != nil {
		return "", err
	}
	return files.names(), nil
}
