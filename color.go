/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

import (
	"github.com/fatih/color"
)

// 输出

func ErrPrintln(v ...interface{}) {
	color.New(color.FgRed, color.Bold).Println(v...)
}

func SuccessPrintln(v ...interface{}) {
	color.New(color.FgGreen).Println(v...)
}

func InfoPrintln(v ...interface{}) {
	color.New(color.FgCyan).Println(v...)
}

func ErrPrintf(f string, v ...interface{}) {
	color.New(color.FgRed, color.Bold).Printf(f, v...)
}

func SuccessPrintf(f string, v ...interface{}) {
	color.New(color.FgGreen).Printf(f, v...)
}

func InfoPrintf(f string, v ...interface{}) {
	color.New(color.FgCyan).Printf(f, v...)
}
