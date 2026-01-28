package commands

import (
	"github.com/urfave/cli/v2"
)

func AllCommands() []*cli.Command {
	commands := []*cli.Command{
		generate,
		Migrate(),
		Worker(),
	}

	return commands
}
