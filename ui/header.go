/*
Project: dirichlet-cli header.go
Created: 2021/12/17 by Landers
*/

package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// 顶部栏

const (
	HeaderTitle = "Dirichlet CLI"
	HeaderText  = "command line tool for Dirichlet by renj.io"
)

func Header() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = HeaderTitle
	p.Text = HeaderText
	p.BorderStyle.Fg = ui.ColorCyan
	p.SetRect(0, 0, HeaderWidth, HeaderHeight)
	p.TextStyle.Fg = ui.ColorGreen

	return p
}
