package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"io/fs"
)

type FileSys struct {
	Fsys   fs.FS
	Config domain.Config
}
