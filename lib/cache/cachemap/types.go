package cachemap

import (
	"context"
	"sync"
	"time"
)

type CacheMap struct {
	ctx    context.Context
	cancel context.CancelFunc
	m      map[string]*Item
	locker sync.RWMutex
}

type Item struct {
	Value       string
	ExpiredTime time.Time
}
