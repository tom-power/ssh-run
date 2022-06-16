package script

import "github.com/tom-power/ssh-run/sshrun/shared"

type Script interface {
	Path(host shared.Host, scriptName string) (string, error)
	Contents(host shared.Host, scriptPath string) (string, error)
	List(host shared.Host) (string, error)
}
