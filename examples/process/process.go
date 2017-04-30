// process is a simple example of spawning a process from the expect package.
package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/google/goexpect"
)

func main() {
	e, _, err := expect.Spawn("echo hello world", 2*time.Second)
	if err != nil {
		glog.Exitf("Shit happended")
	}
	fmt.Println(e)
}
