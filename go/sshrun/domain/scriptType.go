package domain

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
	return keyOr(value, scriptTypes, Default)
}

func (s ScriptType) String() string {
	return scriptTypes[s]
}

func keyOr[V comparable, K comparable](val V, theMap map[K]V, defaultKey K) K {
	for k, v := range theMap {
		if v == val {
			return k
		}
	}
	return defaultKey
}
