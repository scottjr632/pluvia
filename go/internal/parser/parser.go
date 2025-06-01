package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/pluvia/pluvia/utils"
	"gopkg.in/yaml.v3"
)

const suffix = ".pv.yml"

// Config represents the structure of your YAML configuration
type Config struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags,omitempty"`
}

// ParseYAML reads and parses a YAML file into the Config struct
func ParseYAML(filename string) (*Config, error) {
	utils.Invariant(strings.HasSuffix(filename, suffix), "filename must end with .pv.yml")

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	// Create a Config instance to store the parsed data
	config := &Config{}

	// Parse the YAML data into the Config struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}

	return config, nil
}

// WriteYAML writes a Config struct to a YAML file
func WriteYAML(config *Config, filename string) error {
	// Marshal the Config struct to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling YAML: %w", err)
	}

	// Write the YAML data to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing YAML file: %w", err)
	}

	return nil
}
