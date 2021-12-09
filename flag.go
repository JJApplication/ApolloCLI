/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"flag"
)

type fg struct{
	Debug bool
	Addr string
}

const (
	Addr = "/tmp/dirichlet.sock"
)

func flagParse() fg {
	debug := flag.Bool("-d", false, "debug mode")
	addr := flag.String("-a", Addr, "socket addr")
	flag.Parse()

	return fg{
		Debug: *debug,
		Addr: *addr,
	}
}
