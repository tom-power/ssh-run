package domain

import (
	"fmt"
	"strings"

	u "github.com/rjNemo/underscore"
	"github.com/tom-power/ssh-run/sshrun/fp/pred"
)

func (h Host) Ssh() string {
	return fmt.Sprintf("ssh -p %s %s@%s", h.Port, h.User, h.Ip)
}

func (h Host) SshWith(command string, option string) string {
	sshCommandParts := []string{h.Ssh(), option, inDoubleQuotes(withBashLogin(inSingleQuotes(command)))}
	filter := u.Filter(sshCommandParts, pred.IsNotEmpty)
	return strings.Join(filter, " ")
}

func withBashLogin(value string) string {
	return "bash --login -c " + value
}

func inDoubleQuotes(value string) string {
	return "\"" + value + "\""
}

func inSingleQuotes(value string) string {
	return "'" + value + "'"
}
