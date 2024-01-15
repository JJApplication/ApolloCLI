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
	"os"

	"github.com/JJApplication/ApolloCLI/msg"
	"github.com/JJApplication/fushin/errors"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var _version string
var conn string

const (
	_copyright = "Landers1037"
	_powered   = "renj.io"
	_github    = "https://github.com/JJApplication"
)

var CLI = grumble.New(&grumble.Config{
	Name:        ApolloCLI,
	Description: msg.Description,
	Flags: func(f *grumble.Flags) {
		f.String("a", "address", ApolloAddr, msg.FlagAddress)
		f.Bool("v", "version", false, msg.FlagVersion)
		f.Bool("s", "skip", false, msg.FlagVersion)
	},
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
})

func init() {
	CLI.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println(Logo)
	})

	CLI.OnInit(func(a *grumble.App, flags grumble.FlagMap) error {
		address := flags.String("address")
		version := flags.Bool("version")
		skipUds := flags.Bool("skip")
		if version {
			SuccessPrintf("version: %s\n", _version)
			SuccessPrintf("copyright: %s\n", _copyright)
			SuccessPrintf("powered by: %s\n", _powered)
			SuccessPrintf("⇩\n")
			InfoPrintf("github: %s\n", _github)
			a.Close()
			return nil
		}

		conn = ApolloAddr
		if address != ApolloAddr {
			conn = address
			reCreateClient(address)
		}

		// 在某些模式在无需启动uds连接
		if skipUds {
			return nil
		}
		return activeClient()
	})

	CLI.AddCommand(&grumble.Command{
		Name: "version",
		Help: msg.CmdVersion,
		Run: func(c *grumble.Context) error {
			SuccessPrintf("version: %s\n", _version)
			SuccessPrintf("copyright: %s\n", _copyright)
			SuccessPrintf("powered by: %s\n", _powered)
			SuccessPrintf("⇩\n")
			InfoPrintf("github: %s\n", _github)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "create",
		Help: msg.CmdCreate,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgCreate, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			appName := c.Args.String("app")
			if appName == "" {
				return errors.New(msg.ErrCreate)
			}
			InfoPrintf("开始创建微服务: [%s]\n", appName)
			if err := createApp(appName); err != nil {
				SuccessPrintf("微服务[%s]创建完毕\n", appName)
			} else {
				ErrPrintf("微服务[%s]创建失败: %s\n", appName, err.Error())
			}
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "address",
		Help: msg.CmdAddress,
		Run: func(c *grumble.Context) error {
			SuccessPrintf("连接至 %s\n", conn)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "reconnect",
		Help: msg.CmdDial,
		Args: func(a *grumble.Args) {
			a.String("address", msg.ArgAddress, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			addr := c.Args.String("address")
			InfoPrintf("尝试重连至 %s\n", addr)
			if addr == "" {
				return errors.New(msg.ErrNoAddress)
			}
			reCreateClient(addr)
			return activeClient()
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "check",
		Help: msg.CmdCheck,
		Run: func(c *grumble.Context) error {
			if err := ping(); err != nil {
				ErrPrintln(msg.MsgCheckFail)
				return err
			} else {
				SuccessPrintln(msg.MsgCheckSuccess)
			}
			return nil
		},
	})

	// 服务信息
	CLI.AddCommand(&grumble.Command{
		Name: "app",
		Help: msg.CmdApp,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagApp)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				result, err := getAllApp()
				if err != nil {
					return err
				}
				SuccessPrintln(result)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			result, err := getApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(result)
			return nil
		},
	})

	// 服务操作
	CLI.AddCommand(&grumble.Command{
		Name: "start",
		Help: msg.CmdStart,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagStart)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgStart, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				err := startAllApp()
				if err != nil {
					return err
				}
				SuccessPrintln(msg.MsgStartAll)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			err := startApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgStart)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "stop",
		Help: msg.CmdStop,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagStop)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgStop, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				err := stopAllApp()
				if err != nil {
					return err
				}
				SuccessPrintln(msg.MsgStopAll)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			err := stopApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgStop)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "restart",
		Help: msg.CmdRestart,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagRestart)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgRestart, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				err := restartAllApp()
				if err != nil {
					return err
				}
				SuccessPrintln(msg.MsgRestartAll)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			err := restartApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgRestart)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "status",
		Help: msg.CmdStatus,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagStatus)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgStatus, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				result, err := statusAllApp()
				if err != nil {
					SuccessPrintln(result)
					return err
				}
				SuccessPrintln(result)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			result, err := statusApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(result)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "sync",
		Help: msg.CmdSync,
		Flags: func(f *grumble.Flags) {
			f.Bool("a", "all", false, msg.FlagSync)
		},
		Args: func(a *grumble.Args) {
			a.String("name", msg.ArgSync, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.Bool("all") {
				err := syncAllApp()
				if err != nil {
					return err
				}
				SuccessPrintln(msg.MsgSyncAll)
				return nil
			}
			appName := c.Args.String("name")
			if appName == "" {
				return errors.New(msg.ErrNoApp)
			}
			err := syncApp(appName)
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgSync)
			return nil
		},
	})

	// 其他操作
	CLI.AddCommand(&grumble.Command{
		Name: "backup",
		Help: msg.CmdBackup,
		Run: func(c *grumble.Context) error {
			err := backup()
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgBackup)
			return nil
		},
	})

	CLI.AddCommand(&grumble.Command{
		Name: "reload",
		Help: msg.CmdReload,
		Run: func(c *grumble.Context) error {
			err := reload()
			if err != nil {
				return err
			}
			SuccessPrintln(msg.MsgReload)
			return nil
		},
	})

	// NoEngine
	noengineCmds := &grumble.Command{
		Name: "noengine",
		Help: msg.CmdNoEngine,
		Run: func(c *grumble.Context) error {
			fmt.Println(msg.CmdNoEngine)
			return nil
		},
	}
	noengineCmds.AddCommand(&grumble.Command{
		Name: "all",
		Help: msg.CmdNoEngineAll,
		Run: func(c *grumble.Context) error {
			data, err := NoEngineManage("", NoEngineAll)
			if err != nil {
				return err
			}
			res, err := TableApps(data)
			if err != nil {
				return err
			}
			SuccessPrintln(res)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "status",
		Help: msg.CmdNoEngineCheck,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineStatus)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "start",
		Help: msg.CmdNoEngineStart,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineStart)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "stop",
		Help: msg.CmdNoEngineStop,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineStop)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "restart",
		Help: msg.CmdNoEngineRestart,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineReStart)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "pause",
		Help: msg.CmdNoEnginePause,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEnginePause)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "resume",
		Help: msg.CmdNoEngineResume,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineResume)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "remove",
		Help: msg.CmdNoEngineRemove,
		Args: func(a *grumble.Args) {
			a.String("app", msg.ArgNoEngineApp, grumble.Default(""))
		},
		Run: func(c *grumble.Context) error {
			app := c.Args.String("app")
			data, err := NoEngineManage(app, NoEngineRemove)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	noengineCmds.AddCommand(&grumble.Command{
		Name: "refresh",
		Help: msg.CmdNoEngineRefresh,
		Run: func(c *grumble.Context) error {
			data, err := NoEngineManage("", NoEngineRefresh)
			if err != nil {
				return err
			}
			SuccessPrintln(data)
			return nil
		},
	})
	CLI.AddCommand(noengineCmds)

	CLI.SetInterruptHandler(func(a *grumble.App, count int) {
		if count >= 2 {
			InfoPrintln(msg.Thanks)
			os.Exit(0)
		}
		a.Println(msg.Exit)
	})
}
