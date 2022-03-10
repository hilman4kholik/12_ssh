package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"gopkg.in/ini.v1"

	// "io/ioutil"
	"log"
	"os"
)

func main() {
	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
	expath := filepath.Dir(ex)

	fileconfig := expath + "/config.ini"

	cfg, err := ini.Load(fileconfig)
	if err != nil {
		fmt.Println("Faild to read ini file: ", err)
		os.Exit(1)
	}

	
    SSH_ADDRESS  := cfg.Section("settings").Key("address").String() + ":22"
    SSH_USERNAME := cfg.Section("settings").Key("user").String()
    SSH_PASSWORD := cfg.Section("settings").Key("password").String()
	SHELL_SCRIPT := cfg.Section("commands").Key("shellScript").String()
	
    sshConfig := &ssh.ClientConfig{
        User:            SSH_USERNAME,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            ssh.Password(SSH_PASSWORD),
        },
    }

    fmt.Println("auto ssh to: ", cfg.Section("settings").Key("address").String())

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	
	commandshell := strings.Split(SHELL_SCRIPT, ",")
	
	for _, num := range commandshell{
		session, err := client.NewSession()
		if session != nil {
			defer session.Close()
		}
		if err != nil {
			log.Fatal("Failed to create session. " + err.Error())
		}

		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		err = session.Run(num)
		if err != nil {
			log.Fatal("Command execution error. " + err.Error())
		}
		fmt.Println("=============================================")
		fmt.Println("                    DONE                     ")
		fmt.Println("=============================================")
	}

}