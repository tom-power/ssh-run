package domain

type ScriptType int

const (
	Default ScriptType = iota
	Pty
	X11
	Local
)

var scriptTypeToExtension = map[ScriptType]string{
	Default: "",
	Pty:     "pty",
	X11:     "x11",
	Local:   "local",
}

func ScriptTypeFrom(extension string) ScriptType {
	for k, v := range scriptTypeToExtension {
		if v == extension {
			return k
		}
	}
	return Default
}

func (s ScriptType) Extension() string {
	return scriptTypeToExtension[s]
}
