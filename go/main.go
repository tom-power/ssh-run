package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	u "github.com/rjNemo/underscore"
	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/fp"
	"github.com/tom-power/ssh-run/sshrun/fp/pred"
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

	hostName := fp.Stream(getCommands(args)).At(1, "")
	scriptName := fp.Stream(getCommands(args)).At(2, "")

	flags := getFlags(os.Args)

	sshRunFlags := sshrun.RunFlags{
		Help:    u.Any(flags, pred.IsEqualFn("--help")),
		Hosts:   u.Any(flags, pred.IsEqualFn("--hosts")),
		Scripts: u.Any(flags, pred.IsEqualFn("--scripts")),
		Explain: u.Any(flags, pred.IsEqualFn("--explain")),
		Ssh:     u.Any(flags, pred.IsEqualFn("--ssh")),
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
	return u.Filter(args, isNotFlag)
}

func getFlags(args []string) []string {
	return u.Filter(args, isFlag)
}

var flags = []string{"--help", "--explain", "--hosts", "--scripts", "--ssh"}

func isFlag(s string) bool {
	return u.Any(flags, pred.IsEqualFn(s))
}

func isNotFlag(s string) bool {
	return !isFlag(s)
}
