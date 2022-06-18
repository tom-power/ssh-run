package script

import "github.com/tom-power/ssh-run/sshrun/domain"

type Script interface {
	List(host domain.Host) (string, error)
	Path(host domain.Host, scriptName string) (string, error)
	Contents(host domain.Host, scriptName string) (string, error)
	Type(host domain.Host, scriptName string) (string, error)
}
