package handler

import (
	"context"
	"eims/lib/cache"
	"eims/lib/database"
	"eims/lib/log"
	"time"
)

type RunnginCtx struct {
	Logger *log.Logger
	Ctx    context.Context
	DB     *database.Database
	Cache  cache.Cache
	//
}

const (
	sqlExecTimeout = 8 * time.Second
)
