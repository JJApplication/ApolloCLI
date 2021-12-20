/*
Project: dirichlet-cli executer.go
Created: 2021/12/17 by Landers
*/

package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/landers1037/dirichlet_cli/history"
)

// 命令行终端输入

const (
	ExecTitle = "Console"
)

func Exec() *widgets.Paragraph {
	e := widgets.NewParagraph()
	e.Title = ExecTitle
	e.TitleStyle.Fg = ui.ColorGreen
	e.Text = history.CmdList
	e.TextStyle.Fg = ui.ColorGreen

	e.SetRect(0, HeaderHeight + AppHeight, ExecWidth, ExecHeight + HeaderHeight + AppHeight)
	return e
}
