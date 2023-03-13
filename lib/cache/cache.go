package cache

import "time"

type Cache interface {
	Set(key string, value string, expiration time.Duration) error
	SetExpire(key string, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string)
	Close() error
}
