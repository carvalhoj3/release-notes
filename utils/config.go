package utils

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Port     int
	Email    string
	Password string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("utils/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
