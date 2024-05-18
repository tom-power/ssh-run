package domain

import (
	"fmt"
	"strings"

	"github.com/tom-power/ssh-run/sshrun/utils"
	"github.com/tom-power/ssh-run/sshrun/utils/fp"
)

func (h Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", h.Port, h.User, h.Ip)
}

func (h Host) SshWith(command string, option string) string {
	sshCommandParts := []string{h.Ssh(), option, inQuotes(command)}
	filter := fp.Filter(sshCommandParts, utils.IsNotEmpty)
	return strings.Join(filter, " ")
}

func inQuotes(value string) string {
	return "\"" + value + "\""
}
