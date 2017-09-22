// process is a simple example of spawning a process from the expect package.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"regexp"

	"github.com/golang/glog"
	"github.com/google/goexpect"
)

const (
	command = `bash -i`
	endRune = ''
)

func main() {
	e, finCh, err := expect.Spawn(command, -1)
	if err != nil {
		glog.Exit(err)
	}

	usr, err := user.Current()
	if err != nil {
		glog.Exit(err)
	}
	re, err := regexp.Compile(usr.Username)
	if err != nil {
		glog.Exit(err)
	}

	if err := e.Send("cat /etc/passwd\n"); err != nil {
		glog.Exit(err)
	}
	_, _, err = e.Expect(re, -1)
	if err != nil {
		glog.Exit(err)
	}

	go io.Copy(e, os.Stdin)

	var buf = make([]byte, 512)
	for {
		nr, err := e.Read(buf)
		if err != nil {
			glog.Exit(err)
		}
		if bytes.ContainsRune(buf[:nr], endRune) {
			fmt.Println("Found the rune!")
			break
		}
		_, err = os.Stdout.Write(buf[:nr])
		if err != nil {
			glog.Exit(err)
		}
	}

	if err := e.Send("cat /etc/passwd\n"); err != nil {
		glog.Exit(err)
	}
	_, _, err = e.Expect(re, -1)
	if err != nil {
		glog.Exit(err)
	}
	if err := e.Send("exit\n"); err != nil {
		glog.Exit(err)
	}

	fmt.Println(<-finCh)
}
