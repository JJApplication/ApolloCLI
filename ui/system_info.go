/*
Project: dirichlet-cli system_info.go
Created: 2021/12/17 by Landers
*/

package ui

import (
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

// 系统信息

func System() []*widgets.Gauge {
	g := widgets.NewGauge()
	g.Title = "CPU"
	g.Percent = cpuInfo()
	g.SetRect(ExecWidth, HeaderHeight + LoggerHeight, ExecWidth + SystemWidth, HeaderHeight + LoggerHeight + SystemHeight)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorRed

	g2 := widgets.NewGauge()
	g2.Title = "MEMORY"
	g2.Percent = memInfo()
	g2.SetRect(ExecWidth, HeaderHeight + LoggerHeight + SystemHeight, ExecWidth + SystemWidth, HeaderHeight + LoggerHeight + SystemHeight + SystemHeight)
	g2.BarColor = ui.ColorGreen
	g2.BorderStyle.Fg = ui.ColorWhite
	g2.TitleStyle.Fg = ui.ColorRed

	return []*widgets.Gauge{g, g2}
}

// Info 平台 负载
func Info() *widgets.SparklineGroup {
	sp := widgets.NewSparkline()
	sp.Title = "load 1m 5m 15m:"
	sp.TitleStyle.Fg = ui.ColorYellow
	sp.Data = cpuLoad()
	sp.LineColor = ui.ColorCyan
	sp.TitleStyle.Fg = ui.ColorWhite

	slg := widgets.NewSparklineGroup(sp)
	slg.Title = "Sys Load"
	slg.SetRect(ExecWidth, HeaderHeight + LoggerHeight + SystemHeight + SystemHeight, ExecWidth + InfoWidth, HeaderHeight + LoggerHeight + SystemHeight + SystemHeight + InfoHeight)

	return slg
}

func Table() *widgets.Table {
	w := widgets.NewTable()
	w.Rows = [][]string{
		[]string{"Total", "Free", "Avail"},
		memStat(),
	}
	w.Title = "Stat"
	w.TitleStyle.Fg = ui.ColorYellow
	w.SetRect(ExecWidth + InfoWidth, HeaderHeight + LoggerHeight + SystemHeight + SystemHeight, ExecWidth + InfoWidth + InfoWidth, HeaderHeight + LoggerHeight + SystemHeight + SystemHeight + InfoHeight)
	return w
}

func cpuInfo() int {
	data, err := cpu.Percent(time.Second * 0, false)
	if err != nil {
		return 0
	}
	return int(data[0])
}

func cpuLoad() []float64 {
	i, err := load.Avg()
	if err != nil {

	}
	// 扩充
	return []float64{i.Load1, i.Load1, i.Load1, i.Load1, i.Load1,
		i.Load5, i.Load5, i.Load5, i.Load5, i.Load5,
		i.Load15, i.Load15, i.Load15, i.Load15, i.Load15}
}

func memInfo() int {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return int(v.UsedPercent)
}

func memStat() []string {
	v, err := mem.VirtualMemory()
	if err != nil {
		return []string{"", "", ""}
	}
	return []string{
		strconv.FormatUint(v.Total / 1024 / 1024, 10),
		strconv.FormatUint(v.Free / 1024 / 1024, 10),
		strconv.FormatUint(v.Available / 1024 / 1024, 10)}
}