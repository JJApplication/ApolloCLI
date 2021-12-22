/*
Project: dirichlet-cli logger.go
Created: 2021/12/17 by Landers
*/

package ui

import (
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/landers1037/dirichlet_cli/history"
)

// 日志输入

const (
	LoggerTitle = "Logs"
)

func Logger() *widgets.List {
	l := widgets.NewList()
	l.Title = LoggerTitle
	l.TitleStyle.Fg = ui.ColorMagenta
	l.Rows = history.LoggerHist
	l.SetRect(AppWidth, HeaderHeight, LoggerWidth+AppWidth, LoggerHeight+HeaderHeight)

	refresh := func() {
		if len(history.LoggerHist) > 14 {
			l.Rows = history.LoggerHist[len(history.LoggerHist)-14:]
		} else {
			l.Rows = history.LoggerHist
		}
	}

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				refresh()
			}
		}
	}()
	return l
}
