package config

type FcmAPI struct {
	Token string `env:"FCM_API_TOKEN"`
	Send  string `env:"FCM_API_SEND"`
}
