package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		log.Fatal("CONFIG_PATH not found")
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", config)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
