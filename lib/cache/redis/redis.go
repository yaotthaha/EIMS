package redis

import (
	"context"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, options *Options) *Redis {
	r := &Redis{ctx: ctx}
	dialer := &net.Dialer{}
	options.Dialer = func(_ context.Context, network, addr string) (net.Conn, error) {
		return dialer.DialContext(r.ctx, network, addr)
	}
	r.Client = redis.NewClient((*redis.Options)(options))
	return r
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) SetExpire(key string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()
	return r.Client.Expire(ctx, key, expiration).Err()
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()
	result := r.Client.Get(ctx, key)
	return result.Result()
}

func (r *Redis) Del(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()
	r.Client.Del(ctx, key)
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
