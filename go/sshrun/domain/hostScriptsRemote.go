package domain

func (host Host) scriptsRemote() (string, error) {
	hostDir := hostDirWithHome(host.Name, "/home/"+host.User)
	command := "find " + hostDir + " -type f -printf '%f '"
	return runCommandOn(host, command)
}
