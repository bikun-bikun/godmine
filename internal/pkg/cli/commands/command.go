package commands

import (
	"github.com/urfave/cli"
)

type Commands struct {
	cli.Commands
	endpoint string
	apikey   string
}

func NewCommand() (cmd Commands) {

	cmd = Commands{
		[]cli.Command{
			config,
			issue,
		},
		"hoge",
		"huga",
	}

	return cmd
}
