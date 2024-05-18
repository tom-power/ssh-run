package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	. "github.com/tom-power/ssh-run/sshrun/utils/fp"
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

	args := os.Args

	hostName := GetOr(getCommands(args), 1, "")
	scriptName := GetOr(getCommands(args), 2, "")

	flags := getFlags(os.Args)

	sshRunFlags := sshrun.RunFlags{
		Help:    Any(flags, IsEqual("--help")),
		Hosts:   Any(flags, IsEqual("--hosts")),
		Scripts: Any(flags, IsEqual("--scripts")),
		Explain: Any(flags, IsEqual("--explain")),
	}

	sshRun, err := runner.Run(hostName, scriptName, sshRunFlags, getScriptArgs(args))
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
	return Filter(args, isNotFlag)
}

func getScriptArgs(args []string) []string {
	if len(args) > 2 {
		Filter(args[2:], isNotFlag)
	}
	return []string{}
}

func getFlags(args []string) []string {
	return Filter(args, isFlag)
}

var flags = []string{"--help", "--explain", "--hostName", "--scriptName"}

func isFlag(s string) bool {
	return Any(flags, IsEqual(s))
}

func isNotFlag(s string) bool {
	return !isFlag(s)
}
