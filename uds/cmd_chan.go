/*
Project: dirichlet-cli cmd_chan.go
Created: 2021/12/18 by Landers
*/

package uds

import (
	"context"
	"time"
)

// 用于接受cmd通道

const (
	SendTimeOut = "cmd send timeout"
)

func SendCmd(cmd string) string {
	// 防止因uds连接失败导致的卡住
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	cmdChan<-cmd
	// 拿到结果 阻塞的
	var res string
	for {
		select {
		case <-ctx.Done():
			clearChan()
			return SendTimeOut
		case res = <-resChan:
			return res
		}
	}
}

//超时后一定要清空通道
func clearChan() {
	close(cmdChan)
	close(resChan)
}