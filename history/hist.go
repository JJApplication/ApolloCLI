/*
Project: dirichlet-cli hist.go
Created: 2021/12/17 by Landers
*/

package history

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var LoggerHist []string
var CmdList string
var cmdHist [10]string
var seekIndex int = 10
var lock sync.Mutex

const (
	Flag      = "$ "
	Splitter  = "|"
	Space     = "<Space>"
	BackSpace = "<C-<Backspace>>"
	Up        = "<Up>"
	Down      = "<Down>"
	MUp       = "MouseWheelUp"
	MDown     = "MouseWheelDown"
	MLeft     = "<MouseLeft>"
	PageUp    = "<PageUp>"
	PageDown  = "<PageDown>"
)

func init() {
	CmdList = Flag
	cmdHist = [10]string{}
}

func WriteHist(h string) {
	lock.Lock()
	defer lock.Unlock()
	// 增加时间
	LoggerHist = append(LoggerHist, now(h))
}

func CleanHist() {
	lock.Lock()
	defer lock.Unlock()
	LoggerHist = []string{}
}

func now(s string) string {
	var secStr string
	hour, min, sec := time.Now().Clock()
	if sec < 10 {
		secStr = "0" + strconv.Itoa(sec)
	} else {
		secStr = strconv.Itoa(sec)
	}
	return fmt.Sprintf("[%d:%d:%s] %s", hour, min, secStr, s)
}

// WriteCmd cmd列表发送时需要去除>
// 这里是对外显示的cmd以及输出
func WriteCmd(cmd string) {
	lock.Lock()
	defer lock.Unlock()
	switch cmd {
	case BackSpace:
		CmdListBytes := []byte(RealCmd())
		l := len(CmdListBytes)
		if l <= 1 {
			CmdList = Flag
		}
		if l > 1 {
			CmdListBytes = CmdListBytes[:len(CmdListBytes)-1]
			CmdList = Flag + string(CmdListBytes) + Splitter
		}
	case Space:
		CmdList = strings.TrimRight(CmdList, Splitter) + " " + Splitter
	case Up:
		s := SeekLastCmd()
		if s != "" {
			CmdList = s
		}
	case Down:
		s := SeekNextCmd()
		if s != "" {
			CmdList = s
		}
	default:
		CmdList = strings.TrimRight(CmdList, Splitter) + cmd + Splitter
	}
}

func ResultCmd(res string) {
	lock.Lock()
	defer lock.Unlock()

	CmdList = CmdList + "\n" + res + "\n"
}

func CleanCmd() {
	lock.Lock()
	defer lock.Unlock()
	CmdList = Flag
}

func RealCmd() string {
	return strings.TrimLeft(strings.TrimRight(CmdList, Splitter), Flag)
}

func CheckEmptyCmd() bool {
	if CmdList == Flag || CmdList == Flag+Splitter {
		return true
	}
	CmdList = strings.TrimSpace(CmdList)
	if CmdList == Flag || CmdList == Flag+Splitter {
		return true
	}

	return false
}

// 命令历史记录 维持10条

func AddCmdHist(cmd string) {
	for i := range cmdHist {
		if cmdHist[i] == "" {
			cmdHist[i] = cmd
			return
		}
	}
	// 全部遍历后，历史已满，出栈的方式去除元素
	var tmp []string
	tmp = cmdHist[1:]
	for i := range tmp {
		cmdHist[i] = tmp[i]
	}
	cmdHist[9] = cmd
}

func valueCount() int {
	var count int
	for i := range cmdHist {
		if cmdHist[i] != "" {
			count += 1
		}
	}
	return count
}

func SeekLastCmd() string {
	if valueCount() == 0 {
		return ""
	}
	if seekIndex >= valueCount() {
		seekIndex = valueCount() - 1
	}

	if seekIndex < 0 {
		seekIndex = 0
		return ""
	}
	// 向上遍历
	if cmdHist[seekIndex] != "" {
		r := Flag + cmdHist[seekIndex]
		seekIndex -= 1
		return r
	}
	return ""
}

func SeekNextCmd() string {
	if valueCount() == 0 {
		return ""
	}
	if seekIndex >= valueCount() {
		seekIndex = valueCount() - 1
	}

	if seekIndex < 0 {
		seekIndex += 1
		return ""
	}

	if seekIndex >= 10 {
		seekIndex = 10
		return ""
	}
	if cmdHist[seekIndex] != "" {
		r := Flag + cmdHist[seekIndex]
		seekIndex += 1
		return r

	}
	return ""
}
