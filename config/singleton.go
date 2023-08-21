package config

type Singleton struct {
	SingletonUpdateMsec int64 `env:"SINGLETON_UPDATE_MILLIS"`
}
