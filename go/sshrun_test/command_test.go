package sshrun_test

import (
	"strings"
	"testing"

	"github.com/tom-power/ssh-run/sshrun/domain"
)

var testHost = domain.Host{
	Ip:         "192.0.2.1",
	User:       "testUser",
	Name:       "testName",
	Port:       "22",
	PortTunnel: "1081",
}

var commandTypes = []domain.ScriptType{domain.Local, domain.Pty, domain.X11, domain.Default}

func Test_command(t *testing.T) {
	t.Run("can run commands with spaces, multiple lines and escaped characters", func(t *testing.T) {
		for _, commandType := range commandTypes {
			multilineCommand := `
						  multiline \
						  command   
						  `
			script := domain.Script{Type: commandType, Contents: multilineCommand}
			sshRun, _ := testHost.Command(script)

			expected := "multiline command"
			if !strings.Contains(sshRun, expected) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, expected)
			}
		}
	})

	t.Run("can run commands with replacements", func(t *testing.T) {
		for _, commandType := range commandTypes {
			script := domain.Script{Type: commandType, Contents: "$ip$userName$hostName$port$portTunnel$1$2"}
			sshRun, _ := testHost.Command(script)
			if !strings.Contains(sshRun, testHost.Ip) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Ip)
			}
			if !strings.Contains(sshRun, testHost.Name) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Name)
			}
			if !strings.Contains(sshRun, testHost.User) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.User)
			}
			if !strings.Contains(sshRun, testHost.Port) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Port)
			}
			if !strings.Contains(sshRun, testHost.PortTunnel) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.PortTunnel)
			}
		}
	})

}
