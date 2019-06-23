package commands

import (
	"fmt"
	"github.com/bikun-bikun/godmine/internal/pkg/cli/actions"
	"github.com/bikun-bikun/godmine/internal/pkg/cli/flags"
	"github.com/urfave/cli"
)

var issue = cli.Command{
	Name:    "Issue",
	Usage:   "",
	Aliases: []string{"i"},
	Action:  actions.Issue,
	Subcommands: []cli.Command{
		{
			Name:    "template",
			Usage:   "",
			Aliases: []string{"t"},
			Action: func(ctx *cli.Context) {
				fmt.Println("suncommand template: ", ctx.String("profile"))
			},
			Flags: []cli.Flag{
				flags.Profile,
				flags.Create,
			},
		},
	},
	Flags: []cli.Flag{
		flags.Profile,
		flags.Create,
		flags.Project,
		flags.List,
	},
}

var createIssue = cli.Command{
	Name: "Create",
}
