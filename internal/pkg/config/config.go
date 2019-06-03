package config

import (
	"errors"
	"io/ioutil"
	"os"

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
	if !exists(path) {
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
	if exists(path) {
		return errors.New("File Exists")
	}
	c := &Config{ep, k}
	y, err := yaml.Marshal(c)
	if err != nil {
		return errors.New("yaml Marshal error")
	}
	if err = ioutil.WriteFile(path, y, 0644); err != nil {
		return errors.New("Create Config File error")
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getFilePath(profile string) (path string) {
	path = basePath + profile + extension
	return
}
