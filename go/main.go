package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/utils/fp"
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

	hostName := fp.GetOr(getCommands(args), 1, "")
	scriptName := fp.GetOr(getCommands(args), 2, "")

	flags := getFlags(os.Args)

	sshRunFlags := sshrun.RunFlags{
		Help:    fp.Any(flags, fp.IsEqual("--help")),
		Hosts:   fp.Any(flags, fp.IsEqual("--hosts")),
		Scripts: fp.Any(flags, fp.IsEqual("--scripts")),
		Explain: fp.Any(flags, fp.IsEqual("--explain")),
		Ssh:     fp.Any(flags, fp.IsEqual("--ssh")),
	}

	sshRun, err := runner.Run(hostName, scriptName, sshRunFlags)
	if err != nil {
		log.Fatal("sshRun error: ", err)
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
	return fp.Filter(args, isNotFlag)
}

func getFlags(args []string) []string {
	return fp.Filter(args, isFlag)
}

var flags = []string{"--help", "--explain", "--hosts", "--scripts", "--ssh"}

func isFlag(s string) bool {
	return fp.Any(flags, fp.IsEqual(s))
}

func isNotFlag(s string) bool {
	return !isFlag(s)
}
