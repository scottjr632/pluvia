package config

import (
	"fmt"
	"os"

	"github.com/pluvia/pluvia/result"
	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

type Config struct {
	Project struct {
		Name string `yaml:"name"`
	} `yaml:"project"`
	Stack struct {
		Name string `yaml:"name"`
	} `yaml:"stack"`
}

func New() (res result.Result[*Config]) {
	defer result.Recover(&res)

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Failed to read config file: %v\n", err)
		return result.Err[*Config](err)
	}

	config := &Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Failed to parse config file: %v\n", err)
		return result.Err[*Config](err)
	}

	return result.Ok(config)
}
