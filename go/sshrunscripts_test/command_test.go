package sshrunscripts_test

import (
	"strings"
	"testing"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

var testHost = sshrunscripts.Host{"testName", "testUser", "192.0.2.1", "22", "1081"}

var commandTypes = []string{"local", "sudo", "x11", ""}

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
      command \
      \"escaped\"
      `
			sshRun, _ := sshrunscripts.GetCommandSsh(commandType, multilineCommand, testHost, []string{})

			expected := "multiline command \"escaped\""
			if !strings.Contains(sshRun, expected) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, expected)
			}
		}
	})

	t.Run("can run commands with replacements", func(t *testing.T) {
		for _, commandType := range commandTypes {
			sshRun, _ := sshrunscripts.GetCommandSsh(commandType, "$ip$host$user$portSsh$portTunnel$1$2", testHost, []string{"arg1", "arg2"})
			if !strings.Contains(sshRun, testHost.Ip) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Ip)
			}
			if !strings.Contains(sshRun, testHost.Name) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.Name)
			}
			if !strings.Contains(sshRun, testHost.User) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.User)
			}
			if !strings.Contains(sshRun, testHost.PortSsh) {
				t.Errorf("for '%v', '%v' should contain '%v'", commandType, sshRun, testHost.PortSsh)
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
