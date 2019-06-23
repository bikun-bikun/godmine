package commands

import (
	"github.com/bikun-bikun/godmine/internal/pkg/cli/actions"
	"github.com/bikun-bikun/godmine/internal/pkg/cli/flags"
	"github.com/urfave/cli"
)

var config = cli.Command{
	Name:   "Config",
	Action: actions.Config,
	Subcommands: []cli.Command{
		newConfig,
	},
}

var newConfig = cli.Command{
	Name:   "Create",
	Action: actions.NewConfig,
	Flags: []cli.Flag{
		flags.Create,
		flags.Endpoit,
		flags.ApiKey,
		flags.Name,
	},
}
