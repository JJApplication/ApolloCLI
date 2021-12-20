/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package uds

import (
	"flag"
)

var FG Fg

type Fg struct {
	Debug bool
	Addr  string
	Start bool
	Stop bool
}

const (
	Addr = "/tmp/dirichlet.sock"
	ManagerRoot = "/renj.io/app/Dirichlet"
)

func FlagParse() Fg {
	debug := flag.Bool("debug", false, "debug mode")
	addr := flag.String("addr", Addr, "socket addr")
	startServer := flag.Bool("start", false, "start server")
	stopServer := flag.Bool("stop", false, "stop server")
	flag.Parse()

	FG = Fg{
		Debug: *debug,
		Addr:  *addr,
		Start: *startServer,
		Stop: *stopServer,
	}
	return FG
}
