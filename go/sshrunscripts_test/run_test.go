package sshrunscripts_test

import (
	"errors"
	"testing"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

func Test_run(t *testing.T) {
	t.Run("can ssh", func(t *testing.T) {
		actual, _ := testRun("testHostName", "ssh", []string{}, "")

		expected := "ssh -p 22 testUser@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run local", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunLocalScript", []string{}, "")

		if actual != "command" {
			t.Errorf("'%v' should equal '%v'", actual, "command")
		}
	})

	t.Run("can run ssh", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunScript", []string{}, "")

		expected := "ssh -p 22 testUser@192.0.2.1 -f \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with args", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunScriptWithArgs", []string{"arg1", "arg2"}, "")

		expected := "ssh -p 22 testUser@192.0.2.1 -f \"command arg1 arg2\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with sudo", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunSudoScript", []string{}, "")

		expected := "ssh -p 22 testUser@192.0.2.1 -t \"sudo command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh x11", func(t *testing.T) {
		actual, _ := testRun("testHostName", "sshRunX11Script", []string{}, "")

		expected := "ssh -p 22 testUser@192.0.2.1 -Y \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list scripts", func(t *testing.T) {
		actual, _ := testRun("testHostName", "scripts", []string{}, "")

		expected := "echo script anotherScript"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can explain a scripts", func(t *testing.T) {
		actual, _ := testRun("testHostName", "explain", []string{"sshRunScript"}, "")

		expected := "echo ssh -p 22 testUser@192.0.2.1 -f \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list hosts", func(t *testing.T) {
		actual, _ := testRun("hosts", "", []string{}, "")

		expected := "echo testHostName1 testHostName testHostName3"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}

var testGetScripts = func(host sshrunscripts.Host) (string, error) {
	return "script anotherScript", nil
}

func testRun(
	hostName string,
	scriptName string,
	args []string,
	localUserName string) (string, error) {
	return sshrunscripts.Run(
		sshrunscripts.GetHostFromConf([]byte(testConfigText)),
		testGetScriptPath,
		testGetScript,
		sshrunscripts.GetCommandSsh,
		testGetScripts,
		sshrunscripts.GetHostsFromConf([]byte(testConfigText)),
	)(hostName, scriptName, args, localUserName)
}

var testConfigText = `
hosts:
  - name: testHostName1
    user: testUser1
    ip: 192.0.2.1
    portSsh: 22
  - name: testHostName
    user: testUser
    ip: 192.0.2.1
    portSsh: 22
    portTunnel: 1081
  - name: testHostName3
    user: testUser3
    ip: 192.0.2.3
    portSsh: 24`

var testGetScript = func(scriptPath string) (string, error) {
	switch scriptPath {
	case "sshRunLocalScript.local.sh":
		return "command", nil
	case "sshRunScript.sh":
		return "command", nil
	case "sshRunX11Script.sh":
		return "command $1 $2", nil
	case "sshRunScriptWithArgs.sh":
		return "command $1 $2", nil
	case "sshRunSudoScript.sudo.sh":
		return "sudo command", nil
	case "sshRunX11Script.x11.sh":
		return "command", nil
	default:
		return "", errors.New("no script with name " + scriptPath)
	}
}

var testGetScriptPath = func(host sshrunscripts.Host, scriptName string) (string, error) {
	switch scriptName {
	case "ssh":
		return "ssh.ssh.sh", nil
	case "sshRunLocalScript":
		return "sshRunLocalScript.local.sh", nil
	case "sshRunSudoScript":
		return "sshRunSudoScript.sudo.sh", nil
	case "sshRunX11Script":
		return "sshRunX11Script.x11.sh", nil
	default:
		return scriptName + ".sh", nil
	}
}
