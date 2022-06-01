package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/tom-power/ssh-run-scripts/sshrunscripts"
)

func main() {
	conf, err := sshrunscripts.ReadConfig()
	if err != nil {
		log.Println(err)
	}
	var run = sshrunscripts.Run(
		sshrunscripts.GetHostFromConf(conf),
    sshrunscripts.GetScriptPathFromConf,
		sshrunscripts.GetScriptFromConf,
		sshrunscripts.GetCommandSsh,
		sshrunscripts.GetScriptsAll,
		sshrunscripts.GetHostsFromConf(conf),
	)
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	hostName := os.Args[1]
	scriptName := ""
	if len(os.Args) > 2 {
		scriptName = os.Args[2]
	}
	args := []string{""}
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}
	sshRun, err := run(hostName, scriptName, args, usr.Username)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sshRun)
}
