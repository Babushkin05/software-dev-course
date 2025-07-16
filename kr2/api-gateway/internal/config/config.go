package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `yaml:"env"`
	GRPC     GRPCConfig     `yaml:"grpc"`
	HTTP     HTTPConfig     `yaml:"http"`
	Services ServicesConfig `yaml:"services"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type ServicesConfig struct {
	FileStoring  string `yaml:"file_storing"`
	FileAnalysis string `yaml:"file_analysis"`
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
