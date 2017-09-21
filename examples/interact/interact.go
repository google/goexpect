// process is a simple example of spawning a process from the expect package.
package main

import (
	"io"
	"os"
	"regexp"
	"time"

	"github.com/golang/glog"
	"github.com/google/goexpect"
)

const (
	command = `bash -i`
	timeout = 10 * time.Minute
)

var piRE = regexp.MustCompile(`3.14[0-9]*`)

func main() {
	e, _, err := expect.Spawn(command, -1)
	if err != nil {
		glog.Exit(err)
	}

	go io.Copy(e, os.Stdin)
	go io.Copy(os.Stdout, e)

	<-time.After(20 * time.Second)

}
