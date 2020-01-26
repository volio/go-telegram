package config

import "time"

type BotConfig struct {
	Timeout     time.Duration
	EnableProxy bool
	Proxy       string
	Key         string
}
