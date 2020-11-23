/*
VIDEO_1
----------------------
PASOS P/ INSTALACION DE DEPENDENCIAS (en carpeta TP-Go)
	1) go mod init
	2) go get gopkg.in/yaml.v3 (crea el go.mod)
	3) go mod tidy (crea el go.sum)
*/

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig ...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

// Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
