package config

import (
	"errors"
	"github.com/bikun-bikun/godmine/internal/pkg/file"
	"io/ioutil"

	"github.com/go-yaml/yaml"
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
	path := getFilePath(profile)
	if !file.Exists(path) {
		return nil, errors.New("Config not exists!!")
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cnf Config
	if err := yaml.Unmarshal(buf, &cnf); err != nil {
		return nil, err
	}
	return &cnf, nil
}

func NewConfig(profile string, ep string, k string) error {
	path := getFilePath(profile)
	if file.Exists(path) {
		return errors.New("File Exists")
	}

	c := &Config{ep, k}
	y, err := yaml.Marshal(c)
	if err != nil {
		return errors.New("yaml Marshal error")
	}
	if err = file.Save(path, y); err != nil {
		return errors.New("Create Config File error")
	}
	return nil
}

func getFilePath(profile string) (path string) {
	path = basePath + profile + extension
	return
}
