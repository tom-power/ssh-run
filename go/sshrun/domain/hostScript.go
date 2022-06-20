package domain

import (
	"io/fs"
)

func (h Host) Script(fsys fs.FS, scriptName string) (Script, error) {
	scriptContents, err := h.Contents(fsys, scriptName)
	if err != nil {
		return Script{}, err
	}
	scriptType, err := h.Type(fsys, scriptName)
	if err != nil {
		return Script{}, err
	}
	return Script{
		Type:     scriptType,
		Contents: scriptContents,
	}, nil
}
