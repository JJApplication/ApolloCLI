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
	"strings"

	"github.com/JJApplication/ApolloCLI/msg"
	"github.com/JJApplication/fushin/server/uds"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

// 服务启动

func statusApp(name string) (*table.Table, error) {
	tab, err := gotable.Create(name, "状态信息")
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "status",
		Data:      name,
		From:      "",
		To:        nil,
	})
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return tab, fmt.Errorf("%s: %s", msg.ErrStatus, res.Error)
	}
	tab.AddRow([]string{"*", res.Data})
	return tab, nil
}

func statusAllApp() (*table.Table, error) {
	tab, err := gotable.Create("状态信息")
	tab.Align("状态信息", gotable.Left)
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "status-all",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return nil, err
	}

	for _, s := range strings.Split(res.Data, ",") {
		tab.AddRow([]string{s})
	}
	if res.Error != "" {
		return tab, fmt.Errorf("%s: %s", msg.ErrStatusAll, res.Error)
	}
	return tab, nil
}
