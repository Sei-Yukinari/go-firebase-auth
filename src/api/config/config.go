package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type DbConfig struct {
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

const configFile string = "src/api/config/default.toml"

func Load() Config {
	p, _ := os.Getwd()
	var config Config
	fmt.Println(p + configFile)
	if _, err := toml.DecodeFile(p+configFile, &config); err != nil {
		panic(err)
	}
	return config
}
