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

func stopApp(name string) error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "stop",
		Data:      name,
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrStop, res.Error)
	}
	return nil
}

func stopAllApp() error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "stop-all",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrStopAll, res.Error)
	}
	return nil
}
