package config

type Elastic struct {
	Url           string `env:"ELASTIC_URL"`
	ReconnectMsec int    `env:"ELASTIC_RECONNECT_MILLIS"`
}
