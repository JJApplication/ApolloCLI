/*
Project: dirichlet-cli exec.go
Created: 2021/12/11 by Landers
*/

package main

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/landers1037/dirichlet_cli/uds"
)

func start() {
	_, err := exec.Command("bash", "-c", path.Join(uds.ManagerRoot, "start.sh")).Output()
	if err != nil {
		fmt.Printf("server start failed: %s", err.Error())
	} else {
		fmt.Println("server started")
	}
}

func stop() {
	_, err := exec.Command("bash", "-c", path.Join(uds.ManagerRoot, "stop.sh")).Output()
	if err != nil {
		fmt.Printf("server stop failed: %s\n", err.Error())
	} else {
		fmt.Println("server stoped")
	}
}
