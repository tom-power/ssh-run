package main

import (
	"fmt"
	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
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

func getRun() sshrun.Run {
	config, err := config.GetConfigFromFileSystem()
	if err != nil {
		log.Fatal(err)
	}
	return sshrun.GetRun(
		sshrun.GetHostFromConfig(config),
		script.GetScriptPathFromConf,
		script.GetScriptContentsFromHost,
		sshrun.GetCommandSsh,
		script.GetScriptsAll,
		sshrun.GetHostsFromConfig(config),
	)
}
