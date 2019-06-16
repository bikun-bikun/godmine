package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	cf "github.com/bikun-bikun/godmine/internal/pkg/config"
	"github.com/bikun-bikun/godmine/pkg/redmine"

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

var endpoit = cli.StringFlag{
	Name:  "endpoint",
	Value: "",
}

var config = cli.Command{
	Name:    "config",
	Usage:   "",
	Aliases: []string{"c"},
	Flags: []cli.Flag{
		profile,
		create,
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

	cnf, err := cf.Load(ctx.String("profile"))
	if err != nil {
		return errors.New("config load")
	}
	c := redmine.NewClient(cnf.Endpoint, cnf.Apikey)

	ir, err := c.GetIssues("")

	var date []byte

	date, err = json.Marshal(ir)
	fmt.Println(string(date))

	return nil
}

func issueTemplate(ctx *cli.Context) error {

	return nil
}
