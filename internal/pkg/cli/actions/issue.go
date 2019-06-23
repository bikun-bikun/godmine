package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bikun-bikun/godmine/pkg/redmine"
	"github.com/urfave/cli"

	cf "github.com/bikun-bikun/godmine/internal/pkg/config"
)

func Issue(ctx *cli.Context) error {

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

func IssueTemplate(ctx *cli.Context) error {

	return nil
}
