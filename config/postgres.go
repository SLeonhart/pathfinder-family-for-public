package config

import "fmt"

type Postgres struct {
	Host           string `env:"POSTGRES_HOST"`
	User           string `env:"POSTGRES_USER"`
	Password       string `env:"POSTGRES_PASS"`
	DbName         string `env:"POSTGRES_DB_NAME"`
	Port           int    `env:"POSTGRES_PORT"`
	ReconnectMsec  int    `env:"POSTGRES_RECONNECT_MILLIS"`
	MaxConn        int    `env:"POSTGRES_MAX_CONN"`
	MigrationDir   string `env:"POSTGRES_MIGRATION_DIR"`
	MigrationTable string `env:"POSTGRES_MIGRATION_TABLE"`
}

func (p Postgres) DataSource() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.DbName,
	)
}
