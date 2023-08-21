package config

type API struct {
	Debug       bool   `env:"RESTY_DEBUG"`
	Timeout     int    `env:"RESTY_TIMEOUT_MILLIS"`
	Proxy       string `env:"PROXY"`
	FcmAPI      FcmAPI
	TelegramAPI TelegramAPI
	VkAPI       VkAPI
}
