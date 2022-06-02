package sshrunscripts

func scriptPathFromHostRemote(host Host, scriptName string) (string, error) {
	hostDir := hostDir(host.Name, "/home/"+host.User)
	command := "" +
		"cd " + hostDir + "&&" +
		"find . -type f -name '" + scriptName + "*.sh'"
	return runCommandOn(host, command)
}
