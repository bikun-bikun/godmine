package actions

import (
	"errors"
	"github.com/bikun-bikun/godmine/internal/pkg/config"
	"github.com/urfave/cli"
)

func loadConfig() {

}

func Config(ctx *cli.Context) error {

	if ctx.String("create") != "" {

	}
	return nil
}

func NewConfig(ctx *cli.Context) error {
	if ctx.String("endpoint") == "" {
		return errors.New("Endpoint is Required")
	}

	if ctx.String("apikey") == "" {
		return errors.New("ApiKey is Required")
	}

	if err := config.NewConfig(ctx.String("name"), ctx.String("endpoint"), ctx.String("apikey")); err != nil {
		return err
	}
	return nil
}
