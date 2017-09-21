// interact is a simple example of spawning a process from the expect package.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	command = `bc -l`
	timeout = 10 * time.Minute
)

var piRE = regexp.MustCompile(`3.14[0-9]*`)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		glog.Exitf("Usage: process <nr of digits>")
	}

	if err := os.Setenv("BC_LINE_LENGTH", "0"); err != nil {
		glog.Exit(err)
	}

	scale, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		glog.Exit(err)
	}

	e, _, err := expect.Spawn(command, -1)
	if err != nil {
		glog.Exit(err)
	}

	if err := e.Send("scale=" + strconv.Itoa(scale) + "\n"); err != nil {
		glog.Exit(err)
	}

	out, match, err := e.Expect(piRE, timeout)
	if err != nil {
		glog.Exitf("e.Expect(%q,%v) failed: %v, out: %q", piRE.String(), timeout, out)
	}

	fmt.Println(term.Bluef("Pi with %d digits: %s", scale, match[0]))
}

func ioCopy(e *expect.GExpect) {
	go func() {
		if _, err := io.Copy(e, os.Stdin); err != nil {
			glog.Errorf("io.Copy failed: %v", err)
		}
	}()
	go func() {
		if _, err := io.Copy(os.Stdout, e); err != nil {
			glog.Errorf("io.Copy failed: %v,err")
		}
	}()
}
