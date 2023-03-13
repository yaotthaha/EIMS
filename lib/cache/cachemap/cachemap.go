package cachemap

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func (c *CacheMap) autoDel() {
	for {
		select {
		case <-time.After(time.Second):
			c.locker.Lock()
			for k, v := range c.m {
				if !v.ExpiredTime.IsZero() && v.ExpiredTime.Before(time.Now()) {
					delete(c.m, k)
				}
			}
			c.locker.Unlock()
		case <-c.ctx.Done():
			return
		}
	}
}

func New(ctx context.Context) *CacheMap {
	localCtx, cancel := context.WithCancel(ctx)
	c := &CacheMap{
		ctx:    localCtx,
		cancel: cancel,
		m:      make(map[string]*Item),
		locker: sync.RWMutex{},
	}
	go c.autoDel()
	return c
}

func (c *CacheMap) Set(key string, value string, expiration time.Duration) error {
	if key == "" || value == "" {
		return fmt.Errorf("key or value is nil")
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	if expiration <= 0 {
		expiration = 0
	}
	c.m[key] = &Item{
		Value:       value,
		ExpiredTime: time.Now().Add(expiration),
	}
	return nil
}

func (c *CacheMap) SetExpire(key string, expiration time.Duration) error {
	if key == "" {
		return fmt.Errorf("key is nil")
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	item, ok := c.m[key]
	if !ok {
		return fmt.Errorf("key not found")
	}
	if expiration <= 0 {
		item.ExpiredTime = time.Time{}
	} else {
		item.ExpiredTime = time.Now().Add(expiration)
	}
	return nil
}

func (c *CacheMap) Get(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("key is nil")
	}
	c.locker.RLock()
	defer c.locker.RUnlock()
	item, ok := c.m[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	if item.ExpiredTime.Before(time.Now()) {
		return "", fmt.Errorf("key not found")
	}
	return item.Value, nil
}

func (c *CacheMap) Del(key string) {
	if key == "" {
		return
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	delete(c.m, key)
}

func (c *CacheMap) Close() error {
	c.cancel()
	return nil
}
