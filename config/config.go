package config

import (
	"os"
	"path"
)

var (
	BinRoot           = "binRoot"
	CurrentEnv        = "currentEnv"
	InstalledVersions = "installedVersions"
)

type Config struct {
	BinRoot    string `yaml:"binRoot"`
	CurrentEnv string `yaml:"currentEnv"`
}

func New() (*Config, error) {
	c := Config{}
	err := c.setDefaultValues()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) String() string {
	return ""
}

func (c *Config) setDefaultValues() error {
	if c.BinRoot == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		c.BinRoot = path.Join(home, "go")
	}
	return nil
}
