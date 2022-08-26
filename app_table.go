/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

import (
	"strings"

	"github.com/JJApplication/ApolloCLI/msg"
	"github.com/JJApplication/fushin/errors"
	"github.com/JJApplication/fushin/server/uds"
	"github.com/JJApplication/fushin/utils/json"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

// 微服务信息
// 以table表格的形式展示

func getApp(name string) (*table.Table, error) {
	tab, err := gotable.Create(name, "描述信息")
	tab.Align(name, gotable.Left)
	tab.Align("描述信息", gotable.Left)
	if err != nil {
		return nil, err
	}
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "app",
		Data:      name,
		From:      "",
		To:        nil,
	})
	if err != nil {
		return tab, err
	}
	// 序列化数据
	var app App
	err = json.JSON.UnmarshalFromString(res.Data, &app)
	if err != nil {
		return tab, errors.New(msg.ErrApp)
	}
	// 整理数据
	tab.AddRow([]string{"服务ID", app.Meta.ID})
	tab.AddRow([]string{"服务描述", app.Meta.CHSDes})
	tab.AddRow([]string{"服务类型", app.Meta.Type})
	tab.AddRow([]string{"服务状态", app.Meta.ReleaseStatus})

	return tab, nil
}

func getAllApp() (*table.Table, error) {
	tab, err := gotable.Create("微服务")
	tab.Align("微服务", gotable.Left)
	if err != nil {
		return nil, err
	}
	res, err := udsc.SendWithRes(uds.Req{
		Operation: "app-all",
		Data:      "",
		From:      "",
		To:        nil,
	})
	if err != nil {
		return tab, err
	}

	apps := strings.Fields(res.Data)
	if err != nil {
		return tab, errors.New(msg.ErrApp)
	}
	// 整理数据
	for _, app := range apps {
		tab.AddRow([]string{app})
	}

	return tab, nil
}
