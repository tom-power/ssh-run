package domain

func (h Host) scriptsRemote() (string, error) {
	hostDir := hostDirWithHome(h.Name, "/home/"+h.User)
	command := "find " + hostDir + " -type f -printf '%f '"
	return runCommandOn(h, command)
}
