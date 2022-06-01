package sshrunscripts

import (
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func getSession(host Host) (*ssh.Session, error) {
	keyBytes, err := ioutil.ReadFile(homeDir() + "/.ssh/id_rsa")
	if err != nil {
		log.Println("Failed to load private key: ", err)
	}
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	return conn.NewSession()
}
