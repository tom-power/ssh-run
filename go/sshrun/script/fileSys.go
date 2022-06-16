package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

type FileSys struct {
	Fsys   fs.FS
	Config shared.Config
}

func (fsys FileSys) Path(host shared.Host, scriptName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (fsys FileSys) List(host shared.Host) (string, error) {
	//TODO implement me
	panic("implement me")
}
