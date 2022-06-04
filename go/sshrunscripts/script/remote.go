package script

import (
	"bytes"
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/ssh"
)

func getSession(host shared.Host) (*ssh.Session, error) {
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
	conn, err := ssh.Dial("tcp", strings.Join([]string{host.Ip, ":", host.PortSsh}, ""), config)
	if err != nil {
		return nil, err
	}
	return conn.NewSession()
}

func runCommandOn(host shared.Host, command string) (string, error) {
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
