package domain

import (
	"fmt"
	"strings"

	"github.com/tom-power/ssh-run/sshrun/shared"
)

func (h Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", h.Port, h.User, h.Ip)
}

func (h Host) SshWith(command string, option string) string {
	sshCommandParts := []string{h.Ssh(), option, inQuotes(command)}
	filter := shared.Filter(sshCommandParts, shared.IsNotEmpty)
	return strings.Join(filter, " ")
}

func inQuotes(value string) string {
	return "\"" + value + "\""
}
