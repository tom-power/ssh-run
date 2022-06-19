package domain

import (
	"fmt"
)

func (h Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", h.Port, h.User, h.Host)
}

func (h Host) SshWith(command string, option string) string {
	return h.Ssh() + formatOption(option) + formatCommand(command)
}

func formatOption(option string) string {
	if option != "" {
		return fmt.Sprintf(" %s", option)
	}
	return ""
}

func formatCommand(command string) string {
	return fmt.Sprintf(" \"%s\"", command)
}
