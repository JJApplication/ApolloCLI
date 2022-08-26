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

func syncApp(name string) error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "sync",
		Data:      name,
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrSync, res.Error)
	}
	return nil
}

func syncAllApp() error {
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "sync-all",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s: %s", msg.ErrSyncAll, res.Error)
	}
	return nil
}
