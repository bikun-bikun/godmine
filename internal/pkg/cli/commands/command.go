package commands

import (
	"github.com/urfave/cli"
)

func NewCommand() (cmd []cli.Command) {

	cmd = []cli.Command{
		config,
		issue,
	}

	return cmd
}
