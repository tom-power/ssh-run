package sshrunscripts

func scriptsFromHostLocalRemote(host Host) (string, error) {
	hostDir := hostDir(host.Name, "/home/"+host.User)
	command := "find " + hostDir + " -type f -printf '%%f ' | sed \"s/.sh//g\""
	return runCommandOn(host, command)
}
