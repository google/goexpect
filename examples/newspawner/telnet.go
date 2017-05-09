// telnet crates a new Expect sp
package main

import (
	"fmt"
	"net"
	"time"

	expect "github.com/google/goexpect"
	"github.com/google/goterm/term"
	"github.com/ziutek/telnet"
)

func main() {
	fmt.Println(term.Bluef("Telnet spawner example"))

}

func telnetSpawn(network string, addr net.IP, timeout time.Duration) (expect.Expecter, <-chan error, error) {
	conn, err := telnet.Dial(network, addr.String())
	if err != nil {
		return nil, nil, err
	}

	resCh := make(chan error)

	return expect.SpawnGeneric(&expect.GenOptions{
		In:  conn,
		Out: conn,
		Wait: func() error {
			return <-resCh
		},
		Close: func() error {
			close(resCh)
			return conn.Close()
		},
		Check: func() bool { return true },
	}, timeout)
}
