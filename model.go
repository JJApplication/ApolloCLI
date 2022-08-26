/*
Create: 2022/8/26
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

// Package ApolloCLI
package ApolloCLI

// 请求的报文结构
// 最小化的简洁报文

//需要引入octopus-meta

import (
	"github.com/JJApplication/octopus_meta"
)

type App struct {
	Meta octopus_meta.App `json:"meta"`
}
