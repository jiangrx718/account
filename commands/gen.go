package commands

import (
	"account/gopkg/utils"
	"account/model"

	"github.com/urfave/cli/v2"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var generate = &cli.Command{
	Name: "generate",
	Action: func(ctx *cli.Context) error {
		conn := utils.NewMysql("", &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})

		g := gen.NewGenerator(gen.Config{
			OutPath: "internal/dao",
			Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		})

		g.UseDB(conn)

		g.ApplyBasic(
			model.Tags{},
		)

		g.Execute()
		return nil
	},
}
