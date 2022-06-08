package sshrunscripts_test

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"strings"
	"testing"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

var testHost = shared.Host{
	"192.0.2.1",
	"testUser",
	"testName",
	"22",
	"1081",
	false,
}

var commandTypes = []string{"local", "pty", "x11", ""}

func Test_command(t *testing.T) {
	t.Run("can run commandTypes with newlines", func(t *testing.T) {
		sshRun, _ := sshrunscripts.GetCommandSsh("ssh\n", "", testHost, []string{})
		expected := "ssh"
		if !strings.Contains(sshRun, expected) {
			t.Errorf("'%v' should contain '%v'", sshRun, expected)
		}
	})

	t.Run("can run commands with spaces, multiple lines and escaped characters", func(t *testing.T) {
		for _, commandType := range commandTypes {
			multilineCommand := `
      multiline \
      command   
      `
			sshRun, _ := sshrunscripts.GetCommandSsh(commandType, multilineCommand, testHost, []string{})

			expected := "multiline command"
			if !strings.Contains(sshRun, expected) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, expected)
			}
		}
	})

	t.Run("can run commands with replacements", func(t *testing.T) {
		for _, commandType := range commandTypes {
			sshRun, _ := sshrunscripts.GetCommandSsh(commandType, "$host$user$hostName$port$portTunnel$1$2", testHost, []string{"arg1", "arg2"})
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
