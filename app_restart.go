/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

import (
	"fmt"

	"github.com/JJApplication/ApolloCLI/msg"
	"github.com/JJApplication/fushin/server/uds"
)

// 服务启动

func restartApp(name string) error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "restart",
		Data:      name,
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrRestart, res.Error)
	}
	return nil
}

func restartAllApp() error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "restart-all",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrRestartAll, res.Error)
	}
	return nil
}
