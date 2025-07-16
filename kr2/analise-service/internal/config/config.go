package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig struct {
	Port string `yaml:"port"`
}

type StoringServiceConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	GRPC           GRPCConfig           `yaml:"grpc"`
	StoringService StoringServiceConfig `yaml:"storing_service"`
}

func MustLoad() *Config {
	path := FetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exist :" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func FetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
