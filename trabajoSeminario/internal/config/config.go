package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

//...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

//...
func LoadConfig(archivo string) (*Config, error) {
	arch, err := ioutil.ReadFile(archivo)
	if err != nil {
		return nil, err
	}
	var c = &Config{}
	err = yaml.Unmarshal(arch, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
