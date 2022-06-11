package main

import (
	"fmt"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/config"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/script"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"log"
	"os"
)

func main() {
	hostName := shared.SafeSlice(os.Args, 1, "")
	scriptName := shared.SafeSlice(os.Args, 2, "")
	args := []string{""}
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}
	sshRun, err := getRun()(hostName, scriptName, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sshRun)
}

func getRun() sshrunscripts.Run {
	config, err := config.GetConfigFromFileSystem()
	if err != nil {
		log.Fatal(err)
	}
	return sshrunscripts.GetRun(
		sshrunscripts.GetHostFromConfig(config),
		script.GetScriptPathFromConf,
		script.GetScriptContentsFromHost,
		sshrunscripts.GetCommandSsh,
		script.GetScriptsAll,
		sshrunscripts.GetHostsFromConfig(config),
	)
}
