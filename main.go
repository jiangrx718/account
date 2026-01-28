package main

import (
	"account/commands"
	"account/gopkg/log"
	"account/gopkg/utils"
	"context"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/urfave/cli/v2"
)

var configFile string

//go:generate go run main.go generate
func main() {
	app := cli.NewApp()
	app.Version = "local"
	app.Action = commands.Serve
	app.Before = before
	app.After = after
	app.Commands = commands.AllCommands()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Value:       "", //默认从config目录读取
			Usage:       "specify the location of the configuration file",
			Required:    false,
			Destination: &configFile,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Sugar().Fatal(err)
	}
}

// before 初始化相应配置
func before(ctx *cli.Context) error {
	if err := utils.InitViper("account", configFile); err != nil {
		return err
	}

	log.InitFromViper()

	if err := utils.InitOpenTelemetry(context.Background()); err != nil {
		return err
	}

	if err := utils.InitMysqlFromViper(); err != nil {
		return err
	}

	if err := utils.InitSentryFromViper(); err != nil {
		return err
	}

	return nil
}

func after(ctx *cli.Context) error {
	sentry.Flush(2 * time.Second)
	_ = log.Sugar().Sync()
	return nil
}
