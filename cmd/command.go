package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "profile",
		Usage: "Change Profile",
		Value: "default",
	},
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
					Action: func(cxt *cli.Context) error {
						fmt.Println("suncommand template")
						return nil
					},
				},
			},
		},
	}

	return cmd, nil
}

func issue(ctx *cli.Context) error {

	fmt.Printf("main command profile : %+v\n", ctx.GlobalString("profile"))
	return nil
}
