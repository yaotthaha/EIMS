package handler

import (
	"context"
	"eims/lib/cache"
	"eims/lib/database"
	"eims/lib/log"
)

type RunnginCtx struct {
	Logger *log.Logger
	Ctx    context.Context
	DB     *database.Database
	Cache  cache.Cache
	//
}
