package telegram

import "time"

type Config struct {
	Request RequestConfig
	Proxy   ProxyConfig
	Bot     BotConfig
}

type RequestConfig struct {
	LongPollTimeout time.Duration
	RequestTimeout  time.Duration
}

type ProxyConfig struct {
	Enable bool
	Addr   string
}

type BotConfig struct {
	Key string
}
