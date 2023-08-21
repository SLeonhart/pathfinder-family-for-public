package config

type TelegramAPI struct {
	BotToken    string `env:"TELEGRAM_API_BOT_TOKEN"`
	Send        string `env:"TELEGRAM_API_SEND"`
	ChannelId   string `env:"TELEGRAM_API_CHANNEL_ID"`
	AdminChatId string `env:"TELEGRAM_API_ADMIN_CHAT_ID"`
}
