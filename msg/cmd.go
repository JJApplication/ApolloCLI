/*
Create: 2022/8/27
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package msg
package msg

// cmd名称

const (
	CmdVersion = "查看ApolloCLI版本"
	CmdDial    = "重连指定的unix地址"
	CmdAddress = "查看连接的unix地址"
	CmdCheck   = "检查Apollo服务状态"
	CmdApp     = "显示注册的微服务列表"
	CmdStart   = "启动指定微服务"
	CmdStop    = "停止指定微服务"
	CmdRestart = "重启指定微服务"
	CmdStatus  = "查看指定微服务"
	CmdSync    = "同步指定微服务"
	CmdBackup  = "全局同步备份"
	CmdReload  = "重载微服务模型文件"
)

const (
	Description = "Apollo交互式终端"
	Thanks      = "感谢使用ApolloCLI"
	Exit        = "再次按下Ctrl-C退出"
)
