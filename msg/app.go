/*
Create: 2022/8/27
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package msg
package msg

const (
	ErrNoApp      = "未指定微服务名称"
	ErrNoAddress  = "空的unix地址"
	ErrCreate     = "输入的微服务名称为空"
	ErrAppExist   = "微服务已存在"
	ErrApp        = "微服务解析失败"
	ErrStartAll   = "微服务群组启动失败"
	ErrStart      = "微服务启动失败"
	ErrStopAll    = "微服务群组停止失败"
	ErrStop       = "微服务停止失败"
	ErrStatusAll  = "微服务群组检查失败"
	ErrStatus     = "微服务检查失败"
	ErrRestartAll = "微服务群组重启失败"
	ErrRestart    = "微服务重启失败"
	ErrSyncAll    = "微服务群组模型同步失败"
	ErrSync       = "微服务模型同步失败"
	ErrReload     = "微服务模型重载失败"
	ErrBackup     = "全局备份失败"
)

const (
	MsgStartAll   = "微服务群组启动成功"
	MsgStart      = "微服务启动成功"
	MsgStopAll    = "微服务群组停止成功"
	MsgStop       = "微服务停止成功"
	MsgRestartAll = "微服务群组重启成功"
	MsgRestart    = "微服务重启成功"
	MsgSyncAll    = "微服务群组模型同步成功"
	MsgSync       = "微服务模型同步成功"
	MsgBackup     = "全局备份任务已下发"
	MsgReload     = "微服务模型重载成功"
)
