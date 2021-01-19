//package main
//
//import (
//	"github.com/sfreiberg/simplessh"
//	"log"
//)
//
//func SshAndRunCommand() error {
//	var client *simplessh.Client
//	var err error
//
//	// Option A: Using a specific private key path:
//	//if client, err = simplessh.ConnectWithKeyFile("hostname_to_ssh_to", "username", "/home/user/.ssh/id_rsa"); err != nil {
//
//	// Option B: Using your default private key at $HOME/.ssh/id_rsa:
//	//if client, err = simplessh.ConnectWithKeyFile("hostname_to_ssh_to", "username"); err != nil {
//
//	// Option C: Use the current user to ssh and the default private key file:
//	if client, err = simplessh.ConnectWithKeyFile("172.16.120.180", "root", "/Users/fengma/.ssh/id_rsa"); err != nil {
//		return err
//	}
//
//	defer client.Close()
//
//	// Now run the commands on the remote machine:
//	if _, err := client.Exec("ls"); err != nil {
//		log.Println(err)
//	}
//
//	return nil
//}
//func main() {
//
//	err := SshAndRunCommand()
//	if err != nil {
//		panic(err)
//	}
//	//var stdoutBuf, stderrBuf bytes.Buffer
//	//cmd := exec.Command("ssh",  "172.16.120.180",  "ls")
//	//stdoutIn, _ := cmd.StdoutPipe()
//	//stderrIn, _ := cmd.StderrPipe()
//	//var errStdout, errStderr error
//	//stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
//	//stderr := io.MultiWriter(os.Stderr, &stderrBuf)
//	//err := cmd.Start()
//	//if err != nil {
//	//	log.Fatalf("cmd.Start() failed with '%s'\n", err)
//	//}
//	//go func() {
//	//	_, errStdout = io.Copy(stdout, stdoutIn)
//	//}()
//	//go func() {
//	//	_, errStderr = io.Copy(stderr, stderrIn)
//	//}()
//	//err = cmd.Wait()
//	//if err != nil {
//	//	log.Fatalf("cmd.Run() failed with %s\n", err)
//	//}
//	//if errStdout != nil || errStderr != nil {
//	//	log.Fatal("failed to capture stdout or stderr\n")
//	//}
//	//outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
//	//fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
//}

package main

import (
	"fmt"
	"time"
	"github.com/appleboy/easyssh-proxy"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "root",
		Server: "172.16.120.180",
		// Optional key or Password without either we try to contact your agent SOCKET
		//Password: "password",
		// Paste your source content of private key
		// Key: `-----BEGIN RSA PRIVATE KEY-----
		// MIIEpAIBAAKCAQEA4e2D/qPN08pzTac+a8ZmlP1ziJOXk45CynMPtva0rtK/RB26
		// 7XC9wlRna4b3Ln8ew3q1ZcBjXwD4ppbTlmwAfQIaZTGJUgQbdsO9YA==
		// -----END RSA PRIVATE KEY-----
		// `,
		KeyPath: "/Users/fengma/.ssh/id_rsa",
		Port:    "22",
		Timeout: 60 * time.Second,
	}

	// Call Run method with command you want to run on remote server.
	stdout, stderr, done, err := ssh.Run("ls -al", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("don is :", done, "stdout is :", stdout, ";   stderr is :", stderr)
	}

	// kill
	stdout, stderr, done, _ = ssh.Run("ps aux|grep monitor | awk '{print $2}'|xargs kill -9", 60*time.Second)
	fmt.Println("don is :", done, "stdout is :", stdout, ";   stderr is :", stderr)
	//stdout, stderr, done, err = ssh.Run("scp org_riki.sql  root@172.16.120.25:/root/", 60*time.Second)
	//// Handle errors
	//if err != nil {
	//	panic("Can't run remote command: " + err.Error())
	//} else {
	//	fmt.Println("don is :", done, "stdout is :", stdout, ";   stderr is :", stderr)
	//
	//}
}