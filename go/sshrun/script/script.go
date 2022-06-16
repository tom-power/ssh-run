package script

import "github.com/tom-power/ssh-run/sshrun/shared"

type Script interface {
	List(host shared.Host) (string, error)
	Path(host shared.Host, scriptName string) (string, error)
	Contents(host shared.Host, scriptName string) (string, error)
	Type(host shared.Host, scriptName string) (string, error)
}
