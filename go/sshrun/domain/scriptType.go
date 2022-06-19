package domain

import "github.com/tom-power/ssh-run/sshrun/shared"

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

func ParseScriptType(s string) ScriptType {
	return shared.KeyFor(s, scriptTypes, Default)
}

func (scriptType ScriptType) String() string {
	return shared.ValueFor(scriptType, scriptTypes)
}
