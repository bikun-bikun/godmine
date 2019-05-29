package config

const (
	basePath  string = "./"
	extension string = ".yml"
)

type Config struct {
	endpoint string
	apikey   string
}

func (c *Config) Load(profile string) (*Config, error) {
	return nil, nil
}

func getFilePath(profile string) (path string) {
	path = basePath + profile + extension
	return
}
