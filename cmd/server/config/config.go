package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerProt  string `env:"SERVER_PORT" env-default:"6060"`
	NetProtokol string `env:"NET_PROTOKOL" env-default:"tcp"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"info"`
}

var configInstance *Config
var configErr error

func GetConfig() (*Config, error) {
	if configInstance == nil {
		var readConfigOnce sync.Once
		readConfigOnce.Do(func() {
			configInstance = &Config{}
			configErr = cleanenv.ReadEnv(configInstance)
		})
	}

	return configInstance, configErr
}
