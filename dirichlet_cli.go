/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"fmt"
)

var FG fg

func main() {
	FG = flagParse()
	if FG.Start {
		start()
		return
	}

	if FG.Stop {
		stop()
		return
	}

	debug(fmt.Sprintf("[CLI Config]\nDebug: %v\nAddress: %s\n", FG.Debug, FG.Addr))
	SuccessF("Start to Dial...")
	dial()
}
