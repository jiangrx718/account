package commands

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func Migrate() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "数据库迁移",
		Subcommands: []*cli.Command{
			{
				Name:  "up",
				Usage: "迁移失败时，请先手动编辑失败的sql，然后修改schema_migrations表的记录",
				Action: func(ctx *cli.Context) error {
					var n int
					if ctx.NArg() != 0 {
						var nn, err = strconv.ParseInt(ctx.Args().Get(0), 10, 64)
						if nil != err {
							return fmt.Errorf("解析参数失败: %v", err)
						}

						n = int(nn)
					}

					var db, err = sql.Open("mysql", viper.GetString("mysql.dsn"))
					if nil != err {
						return fmt.Errorf("无法连接数据库: %v", err)
					}

					driver, err := mysql.WithInstance(db, &mysql.Config{})
					if nil != err {
						return fmt.Errorf("无法连接数据库: %v", err)
					}

					mi, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", filepath.Join(".", "migrations")), "mysql", driver)
					if nil != err {
						return fmt.Errorf("无法迁移数据库: %v", err)
					}

					if n > 0 {
						if err := mi.Steps(n); nil != err && migrate.ErrNoChange != err {
							return fmt.Errorf("迁移失败,Step %d err:%v", n, err)
						}

						return nil
					}

					if err := mi.Up(); nil != err && migrate.ErrNoChange != err {
						fmt.Printf("迁移失败: %v\n", err)
					}

					version, dirty, err := mi.Version()
					if err == nil {
						fmt.Printf("迁移完成\tVersion:%d dirty:%t\n", version, dirty)
					}
					return err
				},
			},
			{
				Name: "down",
				Action: func(ctx *cli.Context) error {
					var n int
					if 0 != ctx.NArg() {
						var nn, err = strconv.ParseInt(ctx.Args().Get(0), 10, 64)
						if nil != err {
							return fmt.Errorf("解析参数失败: %v", err)
						}

						n = int(nn)
					}

					var db, err = sql.Open("mysql", viper.GetString("mysql.dsn"))
					if nil != err {
						return fmt.Errorf("无法连接数据库: %v", err)
					}

					driver, err := mysql.WithInstance(db, &mysql.Config{})
					if nil != err {
						return fmt.Errorf("无法连接数据库: %v", err)
					}

					mi, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", filepath.Join(".", "migrations")), "mysql", driver)
					if nil != err {
						return fmt.Errorf("无法迁移数据库: %v", err)
					}

					if 0 != n {
						if err := mi.Steps(-n); nil != err && migrate.ErrNoChange != err {
							return fmt.Errorf("迁移失败: %v", err)
						}

						return nil
					}

					if err := mi.Down(); nil != err && migrate.ErrNoChange != err {
						fmt.Printf("迁移失败: %v\n", err)
					}

					return nil
				},
			},
		},
	}
}
