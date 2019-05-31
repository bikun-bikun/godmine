package main

import (
	"encoding/json"
	"fmt"
	"github.com/bikun-bikun/godmine/internal/pkg/config"
	"os"

	"github.com/bikun-bikun/godmine/pkg/redmine"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "gome"
	app.Usage = "This app redmine api kick cli"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		cnf, err := config.Load("default")
		if err != nil {
			fmt.Print(err)
		}
		cl := redmine.NewClient(cnf.Endpoint, cnf.Apikey)
		i, err := cl.GetIssues(1)
		out, _ := json.Marshal(i)
		fmt.Printf(string(out))
		return nil
	}

	app.Run(os.Args)
}
