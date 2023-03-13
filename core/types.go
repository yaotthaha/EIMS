package core

import (
	"context"
	C "eims/config"
	"eims/lib/cache"
	"eims/lib/database"
	"eims/lib/log"
)

type Core struct {
	config C.Config
	logger *log.Logger
	ctx    context.Context
	db     *database.Database
	cache  cache.Cache
}
