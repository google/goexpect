// process is a simple example of spawning a process from the expect package.
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/golang/glog"
	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	command = `zsh -i`
	timeout = 10 * time.Minute
)

var piRE = regexp.MustCompile(`3.14[0-9]*`)

func main() {
	pty, err := term.OpenPTY()
	if err != nil {
		glog.Exit(err)
	}
	backupTerm, _ := term.Attr(os.Stdin)
	myTerm := backupTerm

	myTerm.Raw()
	if err := myTerm.Set(os.Stdin); err != nil {
		glog.Exit(err)
	}
	defer backupTerm.Set(os.Stdin)

	e, _, err := expect.Spawn(command, -1)
	if err != nil {
		glog.Exit(err)
	}

	go io.Copy(pty.Master, os.Stdin)
	go io.Copy(os.Stdout, pty.Master)

	go io.Copy(e, pty.Slave)
	go func() {
		for {
			nr, err := io.Copy(pty.Slave, e)
			if err != nil {
				fmt.Println(term.Redf("Read err: %v", err))
			}
			fmt.Println(term.Greenf("Read nr: %d bytes", nr))
		}
	}()

	<-time.After(20 * time.Second)

}
