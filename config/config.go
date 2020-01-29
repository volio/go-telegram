package config

import "time"

type Config struct {
	LongPollTimeout time.Duration
	RequestTimeout  time.Duration
	EnableProxy     bool
	Proxy           string
	Key             string
}
