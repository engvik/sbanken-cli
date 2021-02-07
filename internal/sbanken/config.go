package sbanken

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AccountAliases map[string]string `yaml:"account-aliases"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &cfg, fmt.Errorf("reading config file: %s, %w", path, err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return &cfg, fmt.Errorf("unmarshal yaml file: %w", err)
	}

	return &cfg, nil
}
