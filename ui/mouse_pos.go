/*
Project: dirichlet-cli mouse_pos.go
Created: 2021/12/20 by Landers
*/

package ui

import (
	"image"

	ui "github.com/gizak/termui/v3"
)

// MouseIn 比较计算鼠标点击事件的位置
func MouseIn(e ui.Event, point image.Point) bool {
	if ev, ok := e.Payload.(ui.Mouse); ok {
		if ev.X <= point.X && ev.Y <= point.Y-HeaderHeight {
			return true
		}
		return false
	}
	return false
}
