package domain

func (Config) scriptsHostRemote(host Host) (string, error) {
	hostDir := hostDirWithHome(host.Name, "/home/"+host.User)
	command := "find " + hostDir + " -type f -printf '%f '"
	return runCommandOn(host, command)
}
