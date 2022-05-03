package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type AppConfig struct {
	Domain   string
	Port     string
	Database string
}

func readConfig() AppConfig {
	dat, err := os.ReadFile("./config.toml")
	if err != nil {
		panic(err)
	}
	var config AppConfig
	if err := toml.Unmarshal(dat, &config); err != nil {
		panic(err)
	}
	return config
}
