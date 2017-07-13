package appconf

import "time"

type Config struct {
	RootDirectory string
	ListenAddress string
	CacheLimit    int64
	DefaultTTL    time.Duration
}
