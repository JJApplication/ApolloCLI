/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

// uds 客户端
// 全局维护

import (
	"sync"

	"github.com/JJApplication/fushin/client/uds"
)

var udsc *uds.UDSClient
var lock sync.Mutex

// 默认使用配置激活
// 在用户指定后重新激活

func init() {
	createClient()
}

func createClient() {
	lock.Lock()
	if udsc == nil {
		udsc = &uds.UDSClient{
			Addr:        ApolloAddr,
			MaxRecvSize: ReadSize,
		}
	}

	lock.Unlock()
}

func reCreateClient(addr string) {
	lock.Lock()
	udsc = &uds.UDSClient{
		Addr:        addr,
		MaxRecvSize: ReadSize,
	}
	lock.Unlock()
}

func activeClient() error {
	return udsc.Dial()
}

func shouldClose() error {
	return udsc.Close()
}
