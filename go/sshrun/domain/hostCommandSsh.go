package domain

import (
	"fmt"
)

func (host Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", host.Port, host.User, host.Host)
}

func (host Host) SshWith(command string, option string) string {
	return host.Ssh() + formatOption(option) + formatCommand(command)
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
