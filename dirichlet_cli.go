/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"fmt"

	"github.com/landers1037/dirichlet_cli/history"
	"github.com/landers1037/dirichlet_cli/uds"
	"github.com/landers1037/dirichlet_cli/ui"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("exited")
			return
		}
	}()

	uds.FlagParse()
	if uds.FG.Start {
		start()
		return
	}

	if uds.FG.Stop {
		stop()
		return
	}

	history.WriteHist("dirichlet cli started...")
	go uds.Dial()
	ui.NewWindow()
}
