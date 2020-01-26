package config

import "time"

type BotConfig struct {
	LongPollTimeout time.Duration
	RequestTimeout  time.Duration
	EnableProxy     bool
	Proxy           string
	Key             string
}
