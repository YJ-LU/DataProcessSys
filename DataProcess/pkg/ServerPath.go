package pkg

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func connect() {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	// 执行远程命令
	cmd := "cd /tmp/;ls -l; tar -czf test.tar.gz hello.txt test.txt;ls -l"
	cmdInfo, err := session.CombinedOutput(cmd)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(string(cmdInfo))
}
