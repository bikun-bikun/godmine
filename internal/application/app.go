package application

import (
	"github.com/bikun-bikun/godmine/internal/pkg/config"
	"github.com/fabulous-tech/go-redmine"
)

type App struct {
	client *redmine.Client
}

func NewApp(config config.Config) *App {
	return &App{client: redmine.NewClient(config.Endpoint, config.Apikey)}
}

func Run() {
}
