package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	config *Config
	mu     sync.Mutex
)

type Config struct {
	API       API
	App       App
	Logger    Logger
	Jaeger    Jaeger
	DB        DB
	Cache     Cache
	Sheduller Sheduller
	Singleton Singleton
	Elastic   Elastic
	Email     Email
}

func Init() {
	mu.Lock()
	defer mu.Unlock()

	_ = godotenv.Load(".env")

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panicf("Load config error: %v", err)
	}

	config = &cfg
}

func Get() *Config {
	if config == nil {
		Init()
	}

	return config
}
