package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

type SSHConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	KeyPath  string
}

// PublicKeyFile parses the SSH key file and returns a ssh.AuthMethod
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return ssh.PublicKeys(key)
}

func (c *SSHConfig) RunCommand(command string) (string, error) {
	var authMethods []ssh.AuthMethod

	// Add password authentication if provided
	if c.Password != "" {
		authMethods = append(authMethods, ssh.Password(c.Password))
	}

	// Add public key authentication if key path is provided
	if c.KeyPath != "" {
		authMethods = append(authMethods, PublicKeyFile(c.KeyPath))
	}
	clientConfig := &ssh.ClientConfig{
		User:            c.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := ssh.Dial("tcp", address, clientConfig)
	if err != nil {
		log.Println("failed to dial address: ", c.Host)
		return "", err
	}

	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(command)
	if err != nil {
		return "", err
	}
	return stdoutBuf.String(), nil
}

func main() {
	config := SSHConfig{
		Host:     "",
		Port:     22,
		Username: "root",
		Password: "",
	}

	command := "ls -l"
	output, err := config.RunCommand(command)
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	fmt.Println("Command output:", output)
}
