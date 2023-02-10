/*
Create: 2023/2/10
Project: ApolloCLI
Github: https://github.com/landers1037
Copyright Renj
*/

package ApolloCLI

import (
	"os"
	"path"

	"github.com/JJApplication/ApolloCLI/msg"
	"github.com/JJApplication/fushin/errors"
	"github.com/JJApplication/fushin/utils/files"
)

var shs = []string{"start.sh", "stop.sh", "check.sh"}

// 创建微服务需要在$APP_ROOT下增加微服务目录 同时增加pig文件的描述
// 当前仅作创建服务 初始化操作脚本的工作
func createApp(appName string) error {
	if files.IsExist(path.Join(APPRoot, appName)) {
		return errors.New(msg.ErrAppExist)
	}
	err := os.Mkdir(path.Join(APPRoot, appName), os.ModeDir)
	if err != nil {
		return err
	}

	// 创建脚本
	var e error
	for _, sh := range shs {
		_, e = files.CreateFile(path.Join(APPRoot, appName, sh), 0755, os.ModeDir)
	}
	if e != nil {
		return e
	}
	return nil
}
