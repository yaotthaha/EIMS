package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	*redis.Client
	ctx context.Context
}

type Options redis.Options

const operateTimeout = 5 * time.Second
