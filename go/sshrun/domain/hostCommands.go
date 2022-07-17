package domain

import (
  "strings"
  "io/fs"
)

func (h Host) Commands(fsys fs.FS) (string, error) {
	scripts, err := h.Scripts(fsys)
  if err != nil {
		return "", err
	}
  return strings.Join(append(sharedCommands(), "ssh", "scripts", "commands", scripts), " "), nil
}
