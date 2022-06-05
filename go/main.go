package main

import (
	"fmt"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/script"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"log"
	"os"
	"os/user"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

func main() {
	hostName := shared.SafeSlice(os.Args, 1, "")
	scriptName := shared.SafeSlice(os.Args, 2, "")
	args := []string{""}
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	sshRun, err := getRun()(hostName, scriptName, args, usr.Username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sshRun)
}

func getRun() sshrunscripts.Run {
	conf, err := sshrunscripts.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return sshrunscripts.GetRun(
		sshrunscripts.GetHostFromConf(conf),
		script.GetScriptPathFromConf,
		script.GetScriptContentsFromHost,
		sshrunscripts.GetCommandSsh,
		script.GetScriptsAll,
		sshrunscripts.GetHostsFromConf(conf),
	)
}
