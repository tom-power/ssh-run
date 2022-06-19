package domain

type ScriptType int

const (
	Default ScriptType = iota
	Ssh
	Pty
	X11
	Local
)

var commandTypes = map[string]ScriptType{"default": Default, "ssh": Ssh, "pty": Pty, "x11": X11, "local": Local}

func ParseCommandType(s string) ScriptType {
	return commandTypes[s]
}

func (c ScriptType) String() string {
	for k, v := range commandTypes {
		if v == c {
			return k
		}
	}
	return ""
}
