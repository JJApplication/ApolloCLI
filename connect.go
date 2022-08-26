/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

import (
	"github.com/JJApplication/fushin/server/uds"
)

// 与Apollo建立连接

func ping() error {
	udsc.Dial()
	if err := udsc.Send(uds.Req{
		Operation: "ping",
		Data:      "",
		From:      "",
		To:        nil,
	}); err != nil {
		return err
	}
	return nil
}
