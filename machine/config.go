package machine

import (
	"time"
)

type Config struct {
	MaxEntryPushConcurrency      int
	MiddlewareAcceptStep         time.Duration
	MaxMiddlewarePushConcurrency int
	EntryCapacity                int
	MiddleCapacity               int
}

func verifyConfig(cfg *Config) *Config {
	if cfg.MaxEntryPushConcurrency <= 0 {
		cfg.MaxEntryPushConcurrency = 1024
	}
	if cfg.MiddlewareAcceptStep.Nanoseconds() == 0 {
		cfg.MiddlewareAcceptStep = time.Millisecond
	}
	if cfg.MaxMiddlewarePushConcurrency <= 0 {
		cfg.MaxMiddlewarePushConcurrency = 1024
	}
	if cfg.EntryCapacity <= 0 {
		cfg.EntryCapacity = 1024
	}
	if cfg.MiddleCapacity <= 0 {
		cfg.MiddleCapacity = 1024
	}

	return cfg
}
