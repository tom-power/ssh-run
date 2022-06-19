package domain

import (
	"bytes"
	"io/ioutil"
	"os/user"
	"strings"

	"golang.org/x/crypto/ssh"
)

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func getSession(host Host) (*ssh.Session, error) {
	keyBytes, err := ioutil.ReadFile(homeDir() + "/.ssh/id_rsa")
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", strings.Join([]string{host.Host, ":", host.Port}, ""), config)
	if err != nil {
		return nil, err
	}
	return conn.NewSession()
}

func runCommandOn(host Host, command string) (string, error) {
	session, err := getSession(host)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run(command); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func scriptsDir(homeDir string) string {
	return homeDir + scriptsPath
}

func hostDirWithHome(hostsName string, homeDir string) string {
	return scriptsDir(homeDir) + "host/" + hostsName + "/"
}
