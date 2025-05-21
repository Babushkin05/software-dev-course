package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `mapstructure:"env"`
	GRPC     GRPCConfig     `mapstructure:"grpc"`
	HTTP     HTTPConfig     `mapstructure:"http"`
	Services ServicesConfig `mapstructure:"services"`
}

type GRPCConfig struct {
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

type ServicesConfig struct {
	FileStoring  string `mapstructure:"file_storing"`
	FileAnalysis string `mapstructure:"file_analysis"`
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
