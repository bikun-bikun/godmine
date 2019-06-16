package main

import (
	"fmt"
	"github.com/bikun-bikun/godmine/cmd"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "godmine"
	app.Usage = "This app redmine api kick cli"
	app.Version = "0.0.1"

	app.Flags = cmd.GlobalFlags
	app.Commands, _ = cmd.NewCommand()

	app.Run(os.Args)
}
