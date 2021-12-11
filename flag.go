/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"flag"
)

type fg struct {
	Debug bool
	Addr  string
	Start bool
	Stop bool
}

const (
	Addr = "/tmp/dirichlet.sock"
	ManagerRoot = "/renj.io/app/Dirichlet"
)

func flagParse() fg {
	debug := flag.Bool("d", false, "debug mode")
	addr := flag.String("a", Addr, "socket addr")
	startDirichelt := flag.Bool("start", false, "start server")
	stopDirichlet := flag.Bool("stop", false, "stop server")
	flag.Parse()

	return fg{
		Debug: *debug,
		Addr:  *addr,
		Start: *startDirichelt,
		Stop: *stopDirichlet,
	}
}
