package actions

import cf "github.com/bikun-bikun/godmine/internal/pkg/config"

type config struct {
	endpoint string
	apikey   string
}

func loadConfig(p string) *config {
	c, _ := cf.Load(p)

	return &config{
		endpoint: c.Endpoint,
		apikey:   c.Apikey,
	}
}
