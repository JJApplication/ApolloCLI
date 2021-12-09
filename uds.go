/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func dial() {
	check()
	c, err := net.Dial("unix", FG.Addr)
	if err != nil {
		fmt.Printf("connect to dirichlet failed: %s\n", err.Error())
		return
	}
	
}

func check() {
	if _, err := os.Stat(FG.Addr); os.IsNotExist(err) {
		fmt.Println(ErrAddr)
	}
}