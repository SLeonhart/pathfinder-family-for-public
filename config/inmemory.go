package config

type InMemory struct {
	ExpirationMsec int64 `env:"MEMORY_CACHE_EXPIRATION_MILLIS"`
}
