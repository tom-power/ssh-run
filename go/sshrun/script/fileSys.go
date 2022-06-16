package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

type FileSys struct {
	Fsys   fs.FS
	Config shared.Config
}
