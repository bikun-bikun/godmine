package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	profile,
}

var profile = cli.StringFlag{
	Name:  "profile, p",
	Usage: "Change Profile",
	Value: "default",
}

var create = cli.StringFlag{
	Name:  "create, c",
	Usage: "Create",
	Value: "",
}

func NewCommand() (cmd []cli.Command, err error) {

	cmd = []cli.Command{
		{
			Name:    "Issue",
			Usage:   "",
			Aliases: []string{"i"},
			Action:  issue,
			Subcommands: []cli.Command{
				{
					Name:    "template",
					Usage:   "",
					Aliases: []string{"t"},
					Action: func(ctx *cli.Context) {
						fmt.Println("suncommand template: ", ctx.String("profile"))
					},
					Flags: []cli.Flag{
						profile,
						create,
					},
				},
			},
			Flags: []cli.Flag{
				profile,
				create,
			},
		},
	}

	return cmd, nil
}

func issue(ctx *cli.Context) error {

	fmt.Printf("main command profile : %+v\n", ctx.String("profile"))
	return nil
}

func issueTemplate(ctx *cli.Context) error {

	return nil
}
