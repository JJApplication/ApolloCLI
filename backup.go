/*
Create: 2022/8/27
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

func backup() error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "backup",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrBackup, res.Data)
	}
	return nil
}
