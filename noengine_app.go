/*
   Create: 2024/1/15
   Project: ApolloCLI
   Github: https://github.com/landers1037
   Copyright Renj
*/

package ApolloCLI

import (
	"encoding/json"
	"fmt"
	"github.com/JJApplication/ApolloCLI/http"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"strings"
)

// 对NoEngine APP操作

const (
	NoEngineAll     = "/api/noengine/all"
	NoEngineStart   = "/api/noengine/x/start"
	NoEngineReStart = "/api/noengine/x/restart"
	NoEngineStop    = "/api/noengine/x/stop"
	NoEnginePause   = "/api/noengine/x/pause"
	NoEngineResume  = "/api/noengine/x/resume"
	NoEngineRemove  = "/api/noengine/x/remove"
	NoEngineStatus  = "/api/noengine/x/status"
	NoEngineRefresh = "/api/noengine/x/refresh"
)

type NoEngineTemplate struct {
	// 域名默认由nginx.conf文件配置不需要单独写在配置文件中
	ServerDomain string           `json:"serverDomain"` // 缺省域名
	ServerName   string           `json:"serverName"`   // 微服务名称
	Volumes      []NoEngineVolume `json:"volumes"`      // 映射卷
	Ports        []NoEnginePort   `json:"ports"`        // 映射端口 开启随机端口时hostPort为随机生成
}

type NoEngineVolume struct {
	HostPath  string `json:"hostPath"`
	InnerPath string `json:"innerPath"`
}

type NoEnginePort struct {
	HostPort  string `json:"hostPort"`
	InnerPort string `json:"innerPort"`
	Proto     string `json:"proto"` // 默认为tcp
}

// NoEngineManage 查询全部服务时app为空
func NoEngineManage(app string, ops string) (interface{}, error) {
	switch ops {
	case NoEngineAll:
		res, err := http.Get(NoEngineAll, nil)
		if err != nil {
			return nil, err
		}
		var result struct {
			Data   map[string]NoEngineTemplate `json:"data"`
			Status string                      `json:"status"`
		}
		_ = json.Unmarshal(res, &result)

		return result.Data, err

	case NoEngineStatus:
		data, err := http.Post(NoEngineStatus, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineStart:
		data, err := http.Post(NoEngineStart, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineStop:
		data, err := http.Post(NoEngineStop, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineReStart:
		data, err := http.Post(NoEngineReStart, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEnginePause:
		data, err := http.Post(NoEnginePause, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineResume:
		data, err := http.Post(NoEngineResume, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineRemove:
		data, err := http.Post(NoEngineRemove, nil, map[string]string{"app": app})
		if err != nil {
			return nil, err
		}
		return string(data), err
	case NoEngineRefresh:
		data, err := http.Post(NoEngineRefresh, nil, nil)
		if err != nil {
			return nil, err
		}
		return string(data), err
	default:
		return nil, nil
	}

}

// TableApps Table表格化输出
func TableApps(d interface{}) (*table.Table, error) {
	data := d.(map[string]NoEngineTemplate)
	tab, err := gotable.Create("NoEngineApps", "Ports", "Volumes")
	tab.Align("NoEngineApps", gotable.Center)
	tab.Align("Ports", gotable.Left)
	tab.Align("Volumes", gotable.Left)
	if err != nil {
		return nil, err
	}
	// 整理数据
	for _, app := range data {
		var ports []string
		for _, port := range app.Ports {
			ports = append(ports, fmt.Sprintf("[%s] %s -> %s", port.Proto, port.HostPort, port.InnerPort))
		}
		portsData := strings.Join(ports, "\n")

		for index, vol := range app.Volumes {
			if index == 0 {
				tab.AddRow([]string{app.ServerName, portsData, fmt.Sprintf("%s -> %s", vol.HostPath, vol.InnerPath)})
			} else {
				tab.AddRow([]string{"", "", fmt.Sprintf("%s -> %s", vol.HostPath, vol.InnerPath)})
			}
		}
	}

	return tab, nil
}
