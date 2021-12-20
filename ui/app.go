/*
Project: dirichlet-cli app.go
Created: 2021/12/17 by Landers
*/

package ui

import (
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/landers1037/dirichlet_cli/uds"
)

// app列表

const (
	AppTitle = "Apps"
)

var data []string

func App() *widgets.List {
	l := widgets.NewList()
	l.Rows = parseAppData()
	l.Title = AppTitle
	l.TitleStyle.Fg = ui.ColorCyan
	l.TextStyle.Fg = ui.ColorYellow
	l.SetRect(0, HeaderHeight, AppWidth, AppHeight + HeaderHeight)

	return l
}

func parseAppData() []string {
	res := uds.SendCmd("app list")
	return strings.Fields(res)
}