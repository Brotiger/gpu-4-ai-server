package config

import (
	"github.com/caarlos0/env/v7"
	"log"
)

type Config struct {
	HTTPAddr   string `env:"HTTP_ADDR" envDefault:":8080"`
	WorkerAddr string `env:"WORKER_ADDR" envDefault:"localhost:50051"`
}

func Load() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}
	return cfg
}

// TODO: Конфигурация для server 