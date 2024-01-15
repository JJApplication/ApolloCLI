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
	CmdVersion         = "查看ApolloCLI版本"
	CmdDial            = "重连指定的unix地址"
	CmdAddress         = "查看连接的unix地址"
	CmdCheck           = "检查Apollo服务状态"
	CmdCreate          = "注册微服务"
	CmdApp             = "显示注册的微服务列表"
	CmdStart           = "启动指定微服务"
	CmdStop            = "停止指定微服务"
	CmdRestart         = "重启指定微服务"
	CmdStatus          = "查看指定微服务"
	CmdSync            = "同步指定微服务"
	CmdBackup          = "全局同步备份"
	CmdReload          = "重载微服务模型文件"
	CmdNoEngine        = "操作NoEngine服务"
	CmdNoEngineAll     = "列举NoEngine服务"
	CmdNoEngineCheck   = "NoEngine服务状态"
	CmdNoEngineStart   = "NoEngine服务启动"
	CmdNoEngineStop    = "NoEngine服务停止"
	CmdNoEngineRestart = "NoEngine服务重启"
	CmdNoEnginePause   = "NoEngine服务暂停"
	CmdNoEngineResume  = "NoEngine服务恢复"
	CmdNoEngineRemove  = "NoEngine服务删除"
	CmdNoEngineRefresh = "NoEngine服务重载"
)

const (
	Description = "Apollo交互式终端"
	Thanks      = "感谢使用ApolloCLI"
	Exit        = "再次按下Ctrl-C退出"
)
