package domain

import utils "github.com/tom-power/ssh-run/sshrun/utils/fp"

type ScriptType int

const (
	Default ScriptType = iota
	Pty
	X11
	Local
)

var scriptTypes = map[ScriptType]string{
	Default: "",
	Pty:     "pty",
	X11:     "x11",
	Local:   "local",
}

func (ScriptType) Parse(value string) ScriptType {
	return utils.KeyOr(value, scriptTypes, Default)
}

func (s ScriptType) String() string {
	return scriptTypes[s]
}
