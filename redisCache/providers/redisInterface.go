package providers

import "time"

type CacheProvider interface {
	Delete(key string) int64
	Set(key string, value interface{}, expiryTime time.Duration) error
	Get(Key string) string
}
