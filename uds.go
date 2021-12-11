/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func dial() {
	check()
	c, err := net.Dial("unix", FG.Addr)
	if err != nil {
		ErrorF("connect to dirichlet failed: %s\n", err.Error())
		return
	}

	for {
		var userInput []string
		fmt.Printf("\x1b[0;%dm%s\x1b[0m", SUCCESS, "> ")
		START_SCAN: userInput, err := readInput()
		debug(userInput...)

		if err != nil && err.Error() == "unexpected newline" {
			goto START_SCAN
		}

		if err != nil {
			ErrorF(ErrScan, err)
			continue
		}

		if exit(userInput...) {
			c.Close()
			ErrorF(ErrExit)
			break
		}

		_, err = c.Write([]byte(strings.Join(userInput, " ")))
		if err != nil {
			ErrorF(err.Error())
			ErrorF(ErrExit)
			break
		}

		var buf = make([]byte, 4096 * 64)
		cnt, err := c.Read(buf)
		if err != nil {
			debug(string(buf[:cnt]))
			c.Close()
			ErrorF(err.Error())
			ErrorF(ErrExit)
			break
		}

		reformatResult(string(buf[:cnt]))
	}
}

func readInput() ([]string, error) {
	var input *bufio.Reader
	input = bufio.NewReader(os.Stdin)
	i, err := input.ReadString('\n')
	if err != nil {
		return []string{}, err
	}
	return strings.Fields(i), nil
}

func check() {
	if _, err := os.Stat(FG.Addr); os.IsNotExist(err) {
		ErrorF(ErrAddr)
	}
}

func exit(s ...string) bool {
	if len(s) <= 0 {
		return false
	}
	if s[0] == "q" || s[0] == "exit" || s[0] == "\\q" {
		return true
	}
	return false
}

func debug(s ...string) {
	if FG.Debug {
		InfoF("[debug] %s", strings.Join(s, " "))
	}
}

func reformatResult(s string) {
	if strings.Contains(s, "error") || strings.Contains(s, "err") ||
		strings.Contains(s, "failed") || strings.Contains(s, "incorrect") ||
		strings.Contains(s, "cmd not support") {
		ErrorF(s)
	} else {
		SuccessF(s)
	}
}