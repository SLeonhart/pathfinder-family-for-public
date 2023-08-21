package config

import (
	"fmt"
)

type Jaeger struct {
	Host string `env:"JAEGER_HOST"`
	Port int    `env:"JAEGER_PORT"`
}

func (j Jaeger) ServiceName(cfg *Config) string {
	return fmt.Sprintf("%s_%s", cfg.App.ServiceName, cfg.App.Mode)
}

func (j Jaeger) Server() string {
	return fmt.Sprintf("%s:%d", j.Host, j.Port)
}
