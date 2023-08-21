package config

import (
	"fmt"
)

type App struct {
	// Уникальное название сервиса. Может использоваться в Редисе для префикса, в Jaeger и т.п.
	ServiceName string `env:"SERVICE_NAME"`
	// Возможные значения: debug, test, release
	Mode string `env:"APP_MODE"`
	// Порт, на котором слушает gin
	Port    int    `env:"APP_PORT"`
	Version string `env:"APP_VERSION"`
	HostUrl string `env:"HOST_URL"`
}

func (a App) ServerAddress() string {
	return fmt.Sprintf(":%d", a.Port)
}
