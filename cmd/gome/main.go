package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "gome"
	app.Usage = "This app redmine api kick cli"
	app.Version = "0.0.1"

	app.Run(os.Args)
}
