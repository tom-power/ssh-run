package sshrun_test

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_run(t *testing.T) {
	t.Run("can ssh", func(t *testing.T) {
		actual, _ := testRun("testHostName", "ssh", []string{}, testConfig)
		expected := "ssh -p 22 testUser@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can ssh on no command", func(t *testing.T) {
		actual, _ := testRun("testHostName", "", []string{}, testConfig)
		expected := "ssh -p 22 testUser@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run local", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunLocalScript", []string{}, testConfig)
		if actual != "command" {
			t.Errorf("'%v' should equal '%v'", actual, "command")
		}
	})

	t.Run("can run ssh", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshruncript", []string{}, testConfig)
		expected := "ssh -p 22 testUser@192.0.2.1 \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with args", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshruncriptWithArgs", []string{"arg1", "arg2"}, testConfig)
		expected := "ssh -p 22 testUser@192.0.2.1 \"command arg1 arg2\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with pty", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunPtyScript", []string{}, testConfig)
		expected := "ssh -p 22 testUser@192.0.2.1 -t \"pty command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh x11", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunX11Script", []string{}, testConfig)

		expected := "ssh -p 22 testUser@192.0.2.1 -Y \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list scripts", func(t *testing.T) {
		actual, _ := testRun("testHostName", "scripts", []string{}, testConfig)
		expected := "echo script anotherScript"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list hosts", func(t *testing.T) {
		actual, _ := testRun("hosts", "", []string{}, testConfig)
		expected := "echo testHostName1 testHostName testHostName3"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can explain host", func(t *testing.T) {
		actual, _ := testRun("testHostName", "explain", []string{}, testConfig)
		expected := `Host{Host:"192.0.2.1", User:"testUser", Name:"testHostName", Port:"22", PortTunnel:"1081"}`
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}

func testRun(
	hostName string,
	scriptName string,
	args []string,
	config shared.Config,
) (string, error) {
	return sshrun.GetRun(sshrun.GetCommandSsh, config, testScript)(hostName, scriptName, args)
}

var testHosts = []shared.Host{
	{
		Name: "testHostName1",
		User: "testUser1",
		Host: "192.0.2.1",
		Port: "22",
	},
	{
		Name:       "testHostName",
		User:       "testUser",
		Host:       "192.0.2.1",
		Port:       "22",
		PortTunnel: "1081",
	},
	{
		Name: "testHostName3",
		User: "testUser3",
		Host: "192.0.2.1",
		Port: "22",
	},
}

var testConfig = shared.Config{
	Hosts: testHosts,
}

type TestScript struct{}

func (ts TestScript) Contents(host shared.Host, scriptPath string) (string, error) {
	switch scriptPath {
	case "sshRunLocalScript.local.sh":
		return "command", nil
	case "sshruncript.sh":
		return "command", nil
	case "sshRunX11Script.sh":
		return "command $1 $2", nil
	case "sshruncriptWithArgs.sh":
		return "command $1 $2", nil
	case "sshRunPtyScript.pty.sh":
		return "pty command", nil
	case "sshRunX11Script.x11.sh":
		return "command", nil
	default:
		return "", errors.New("no script with name " + scriptPath)
	}
}

func (ts TestScript) Path(host shared.Host, scriptName string) (string, error) {
	switch scriptName {
	case "ssh":
		return "ssh.ssh.sh", nil
	case "sshRunLocalScript":
		return "sshRunLocalScript.local.sh", nil
	case "sshRunPtyScript":
		return "sshRunPtyScript.pty.sh", nil
	case "sshRunX11Script":
		return "sshRunX11Script.x11.sh", nil
	default:
		return scriptName + ".sh", nil
	}
}

func (ts TestScript) List(host shared.Host) (string, error) {
	return "script anotherScript", nil
}

var testScript = TestScript{}

var testGetScript = func(host shared.Host, scriptPath string, config shared.Config) (string, error) {
	switch scriptPath {
	case "sshRunLocalScript.local.sh":
		return "command", nil
	case "sshruncript.sh":
		return "command", nil
	case "sshRunX11Script.sh":
		return "command $1 $2", nil
	case "sshruncriptWithArgs.sh":
		return "command $1 $2", nil
	case "sshRunPtyScript.pty.sh":
		return "pty command", nil
	case "sshRunX11Script.x11.sh":
		return "command", nil
	default:
		return "", errors.New("no script with name " + scriptPath)
	}
}
