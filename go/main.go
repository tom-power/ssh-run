package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func main() {
	homeDirFs, err := getHomeDirFs()
	if err != nil {
		log.Fatal(err)
	}
	config, err := getConfigFs(homeDirFs).GetConfig()
	if err != nil {
		log.Fatal("config error:", err)
	}
	runner := sshrun.Runner{
		Config: config,
		Fsys:   homeDirFs,
	}

	hostName := shared.GetOr(getCommands(os.Args), 1, "")
	scriptName := shared.GetOr(getCommands(os.Args), 2, "")

	helpFlag := flag.Bool("help", false, "help")
	hostsFlag := flag.Bool("hosts", false, "list hosts")
	scriptsFlag := flag.Bool("scripts", false, "list scripts for host")
	explainFlag := flag.Bool("explain", false, "explain host or script")
	scriptArgsFlag := flag.String("scriptArgs", "", "arguments to pass to script")

	flag.Parse()

	flags := sshrun.RunFlags{
		Help:       *helpFlag,
		Hosts:      *hostsFlag,
		Scripts:    *scriptsFlag,
		Explain:    *explainFlag,
		ScriptArgs: getScriptArgs(*scriptArgsFlag),
	}

	sshRun, err := runner.Run(hostName, scriptName, flags)
	if err != nil {
		log.Fatal("runner error:", err)
	}
	fmt.Println(sshRun)
}

func getConfigFs(homeDirFs fs.FS) config.ConfigFs {
	return config.ConfigFs{
		Fsys:      homeDirFs,
		ConfigDir: ".config/ssh-run",
		SshPath:   ".ssh/config",
	}
}

func getHomeDirFs() (fs.FS, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.DirFS(homeDir), nil
}

func getCommands(args []string) []string {
	return shared.Filter(args, isCommand)
}

func isCommand(s string) bool {
	return !isArgs(s)
}

func isArgs(s string) bool {
	return strings.HasPrefix(s, "--") || strings.HasPrefix(s, "-")
}

func getScriptArgs(scriptArgs string) []string {
	return strings.Split(scriptArgs, ",")
}
