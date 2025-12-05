package commands

import (
	"account/commands/generate"
	"account/commands/gorm"
	"account/commands/migrate"
	"account/commands/worker"

	"github.com/urfave/cli/v2"
)

func All() []*cli.Command {
	commands := []*cli.Command{
		migrate.Command(),
		generate.Command(),
		gorm.Command(),
		worker.Command(),
	}
	return commands
}
