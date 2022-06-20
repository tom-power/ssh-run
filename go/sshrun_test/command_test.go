package sshrun_test

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"strings"
	"testing"
)

var testHost = domain.Host{
	Host:       "192.0.2.1",
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
			sshRun, _ := testHost.Command(script, []string{})

			expected := "multiline command"
			if !strings.Contains(sshRun, expected) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, expected)
			}
		}
	})

	t.Run("can run commands with replacements", func(t *testing.T) {
		for _, commandType := range commandTypes {
			script := domain.Script{Type: commandType, Contents: "$host$user$hostName$port$portTunnel$1$2"}
			sshRun, _ := testHost.Command(script, []string{"arg1", "arg2"})
			if !strings.Contains(sshRun, testHost.Host) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Host)
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
			if !strings.Contains(sshRun, "arg1") {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, "arg1")
			}
			if !strings.Contains(sshRun, "arg2") {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, "arg2")
			}
		}
	})

}
