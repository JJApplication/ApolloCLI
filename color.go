/*
Project: dirichlet-cli color.go
Created: 2021/12/9 by Landers
*/

package main

import (
	"fmt"
)

const (
	ERROR   = 31
	SUCCESS = 32
	INFO    = 36
	PANIC   = 35
)

const Fmt = "\x1b[0;%dm%s\x1b[0m\n"

func ErrorF(f string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(Fmt, ERROR, fmt.Sprintf(f, args...))
	} else {
		fmt.Printf(Fmt, ERROR, f)
	}
}

func InfoF(f string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(Fmt, INFO, fmt.Sprintf(f, args...))
	} else {
		fmt.Printf(Fmt, INFO, f)
	}
}

func SuccessF(f string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(Fmt, SUCCESS, fmt.Sprintf(f, args...))
	} else {
		fmt.Printf(Fmt, SUCCESS, f)
	}
}
