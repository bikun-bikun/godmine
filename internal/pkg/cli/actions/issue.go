package actions

import (
	"encoding/json"
	"fmt"
	"github.com/bikun-bikun/godmine/pkg/redmine"
	"github.com/urfave/cli"

	cf "github.com/bikun-bikun/godmine/internal/pkg/config"
)

func Issue(ctx *cli.Context) error {

	cnf, err := cf.Load(ctx.String("profile"))
	if err != nil {
		return err
	}
	c := redmine.NewClient(cnf.Endpoint, cnf.Apikey)

	ir, err := c.GetIssues(ctx.String("project"))

	var date []byte

	if ctx.Bool("list") {
		var in []redmine.IdName
		for _, i := range ir {
			in = append(in, redmine.IdName{ID: i.Id, Name: i.Subject})
		}
		date, err = json.Marshal(in)
		fmt.Println(string(date))
		return nil
	}
	date, err = json.Marshal(ir)
	fmt.Println(string(date))

	return nil
}

func IssueTemplate(ctx *cli.Context) error {

	return nil
}

func CreateIssue(ctx *cli.Context) error {

	return nil
}
