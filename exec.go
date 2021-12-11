/*
Project: dirichlet-cli exec.go
Created: 2021/12/11 by Landers
*/

package main

import (
	"fmt"
	"os/exec"
	"path"
)

func start() {
	_, err := exec.Command("bash", "-c", path.Join(ManagerRoot, "start.sh")).Output()
	if err != nil {
		fmt.Printf("server start failed: %s\n", err.Error())
	} else {
		fmt.Printf("server started")
	}
}

func stop() {
	_, err := exec.Command("bash", "-c", path.Join(ManagerRoot, "stop.sh")).Output()
	if err != nil {
		fmt.Printf("server stop failed: %s\n", err.Error())
	} else {
		fmt.Printf("server stoped")
	}
}
