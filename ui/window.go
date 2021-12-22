/*
Project: dirichlet-cli window.go
Created: 2021/12/15 by Landers
*/

package ui

import (
	"fmt"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/landers1037/dirichlet_cli/history"
	"github.com/landers1037/dirichlet_cli/uds"
)

// NewWindow 窗口类
func NewWindow() {
	if err := ui.Init(); err != nil {
		history.WriteHist(fmt.Sprintf("dirichlet ui init failed: %s", err.Error()))
	}
	defer ui.Close()

	header := Header()
	app := App()
	logger := Logger()
	exec := Exec()
	system := System()
	info := Info()
	table := Table()

	render := func() {
		ui.Render(header, app, logger, exec, system[0], system[1], info, table)
	}

	refresh := func() {
		system[0].Percent = cpuInfo()
		system[1].Percent = memInfo()
		info.Sparklines[0].Data = cpuLoad()
		table.Rows[1] = memStat()
		render()
	}

	render()
	events := ui.PollEvents()
	ticker := time.NewTicker(time.Second * 1)
	var mouseInApp bool
	for {
		select {
		case e := <-events:
			// 处理鼠标点击事件获取选中的区域
			if e.ID == history.MLeft && MouseIn(e, app.Size()) {
				mouseInApp = true
				app.BorderStyle.Fg = ui.ColorCyan
			} else if e.ID == history.MLeft && !MouseIn(e, app.Size()) {
				mouseInApp = false
				app.BorderStyle.Fg = ui.ColorWhite
			}

			// 选中区域为app时支持上下滚动
			if mouseInApp == true {
				if e.ID == history.MUp || e.ID == history.Up {
					app.ScrollUp()
					app.Size()
					refresh()
				} else if e.ID == history.MDown || e.ID == history.Down {
					app.ScrollDown()
					refresh()
				}
			}

			// 监听翻页事件，支持终端输入的结果滚动
			if e.ID == history.PageUp {
				exec.ScrollUp()
				refresh()
			} else if e.ID == history.PageDown {
				exec.ScrollDown()
				refresh()
			}

			// 监听常规的键盘事件
			if e.Type == ui.KeyboardEvent {
				// 监听退出事件
				if e.ID == Esc || e.ID == CtrlC {
					uds.DeferExit()
					return
				}
				// 监听回车输入事件 拼接输入
				if e.ID != Enter && IsKeys(e.ID) {
					history.WriteCmd(e.ID)
					exec.Rows = strings.Split(history.CmdList, "\n")
					refresh()
				}
				if e.ID == Enter {
					if history.CheckEmptyCmd() {
						history.CleanCmd()
						history.WriteCmd("type [show] for more?")
						exec.Rows = strings.Split(history.CmdList, "\n")
						history.CleanCmd()
						if len(history.LoggerHist) > 14 {
							logger.Rows = history.LoggerHist[len(history.LoggerHist)-14:]
						} else {
							logger.Rows = history.LoggerHist
						}
						refresh()
					} else {
						history.AddCmdHist(history.RealCmd())
						history.WriteHist(fmt.Sprintf("[cmd received] %s", history.RealCmd()))
						res := uds.SendCmd(history.RealCmd(), 5)
						history.ResultCmd(res)
						exec.Rows = strings.Split(history.CmdList, "\n")
						history.CleanCmd()
						if len(history.LoggerHist) > 14 {
							logger.Rows = history.LoggerHist[len(history.LoggerHist)-14:]
						} else {
							logger.Rows = history.LoggerHist
						}
						refresh()
					}
				}
			}
		case <-ticker.C:
			refresh()
		}
	}
}
