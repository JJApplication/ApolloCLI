/*
Project: dirichlet-cli keys.go
Created: 2021/12/18 by Landers
*/

package ui

import (
	"github.com/landers1037/dirichlet_cli/history"
)

var keys = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	history.Space, history.BackSpace, history.Up, history.Down}

const (
	Enter = "<Enter>"
	Esc = "<Esc>"
	CtrlC = "<C-c>"
)

func IsKeys(k string) bool {
	for _, i := range keys {
		if k == i {
			return true
		}
	}
	return false
}
