package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

type KafkaConfig struct {
	Broker  string `yaml:"broker"`
	Topic   string `yaml:"topic"`
	GroupID string `yaml:"group_id"`
}

type Config struct {
	GRPC     GRPCConfig     `yaml:"grpc"`
	Postgres DatabaseConfig `yaml:"postgres"`
	Kafka    KafkaConfig    `yaml:"kafka"`
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
