// process is a simple example of spawning a process from the expect package.
// This one changes the logging package to the standard one and also logs to stdout.
package main

import (
	"flag"
	"fmt"
	"log"
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

type StdOutLogger struct {
	lr *log.Logger
}

func (l *StdOutLogger) Info(args ...interface{}) {
	fmt.Println(args...)
	l.lr.Print(args...)
}

func (l *StdOutLogger) Infof(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
	l.lr.Printf(format, args...)
}

func (l *StdOutLogger) Warning(args ...interface{}) {
	fmt.Println(args...)
	l.lr.Print(args...)
}

func (l *StdOutLogger) Warningf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
	l.lr.Printf(format, args...)
}

func (l *StdOutLogger) Error(args ...interface{}) {
	fmt.Println(args...)
	l.lr.Print(args...)
}

func (l *StdOutLogger) Errorf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
	l.lr.Printf(format, args...)
}

var piRE = regexp.MustCompile(`3.14[0-9]*`)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		glog.Exitf("Usage: process <nr of digits>")
	}

	f, err := os.Create("/tmp/le_log.txt")
	if err != nil {
		os.Exit(10)
	}
	defer f.Close()

	expect.ChangeLogger(&StdOutLogger{
		lr: log.New(f, "Prefix", log.LstdFlags),
	})

	if err := os.Setenv("BC_LINE_LENGTH", "0"); err != nil {
		os.Exit(1)
	}

	scale, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		os.Exit(2)
	}

	e, ch, err := expect.Spawn(command, -1, expect.Verbose(true))
	if err != nil {
		os.Exit(3)
	}

	if err := e.Send("scale=" + strconv.Itoa(scale) + "\n"); err != nil {
		os.Exit(4)
	}
	if err := e.Send("4*a(1)\n"); err != nil {
		os.Exit(5)
	}
	_, match, err := e.Expect(piRE, timeout)
	if err != nil {
		os.Exit(6)
	}
	if err := e.Send("quit\n"); err != nil {
		os.Exit(7)
	}
	<-ch
	os.Stdout.Sync()
	f.Sync()

	fmt.Println(term.Bluef("Pi with %d digits: %s", scale, match[0]))
}
