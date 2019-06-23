package main

import (
	"github.com/urfave/cli"
	"os"

	"github.com/bikun-bikun/godmine/internal/pkg/cli/commands"
)

func main() {
	app := cli.NewApp()

	app.Name = "godmine"
	app.Usage = "This app redmine api kick cli"
	app.Version = "0.0.1"

	app.Commands = commands.NewCommand()

	app.Run(os.Args)
}
