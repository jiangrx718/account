package commands

import (
	"github.com/urfave/cli/v2"
)

// Worker go run main.go worker start
func Worker() *cli.Command {
	return &cli.Command{
		Name: "worker",
		Subcommands: []*cli.Command{
			{
				Name: "start",
				Action: func(ctx *cli.Context) error {
					//dao.SetDefault(utils.DefaultMysql())
					//graceful.Start(sea_tunnel.InitWorkerSeaTunnel()) // 回调SeaTunnel任务执行结果
					//graceful.Wait()
					return nil
				},
			},
		},
	}
}
