package ssh_test

import (
	"log"
	"os"
	"testing"

	"golang.org/x/crypto/ssh"
)

var username = "root"
var password = "Xss8271329"
var host = "114.116.230.93:22"

func TestPassword(t *testing.T) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server:", err)
	}
	log.Println(string(client.ClientVersion()))
}

func TestExecute(t *testing.T) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server:", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session:", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	var commandToExecute = "hostname&&ifconfig"
	session.Run(commandToExecute)
	if err != nil {
		log.Fatal("Error command execute :", err)
	}
}
