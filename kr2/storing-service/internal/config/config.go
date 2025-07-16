package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string     `yaml:"env"`
	GRPC GRPCConfig `yaml:"grpc"`
	S3   S3Config   `yaml:"s3"`
	PG   PGConfig   `yaml:"postgres"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type S3Config struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
	UseSSL    bool   `yaml:"use_ssl"`
}

type PGConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func MustLoad() *Config {
	path := FetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	// Override sensitive data with env vars if available
	cfg.PG.User = fallback(cfg.PG.User, os.Getenv("POSTGRES_USER"))
	cfg.PG.Password = fallback(cfg.PG.Password, os.Getenv("POSTGRES_PASSWORD"))
	cfg.PG.DBName = fallback(cfg.PG.DBName, os.Getenv("POSTGRES_DB"))

	cfg.S3.AccessKey = fallback(cfg.S3.AccessKey, os.Getenv("S3_ACCESS_KEY"))
	cfg.S3.SecretKey = fallback(cfg.S3.SecretKey, os.Getenv("S3_SECRET_KEY"))
	cfg.S3.Bucket = fallback(cfg.S3.Bucket, os.Getenv("S3_BUCKET"))

	return &cfg
}

func fallback(value, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
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
