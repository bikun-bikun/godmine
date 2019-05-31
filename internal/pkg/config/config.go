package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

const (
	basePath  string = "./config/"
	extension string = ".yml"
)

type Config struct {
	Endpoint string `yaml:"endpoint"`
	Apikey   string `yaml:"apikey"`
}

func Load(profile string) (*Config, error) {
	buf, err := ioutil.ReadFile(getFilePath(profile))
	if err != nil {
		return nil, err
	}

	var cnf Config
	if err := yaml.Unmarshal(buf, &cnf); err != nil {
		return nil, err
	}
	return &cnf, nil
}

func getFilePath(profile string) (path string) {
	path = basePath + profile + extension
	return
}
