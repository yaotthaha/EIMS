package core

import (
	"context"
	C "eims/config"
	"eims/handler"
	"eims/lib/cache/cachemap"
	"eims/lib/cache/redis"
	"eims/lib/database"
	"eims/lib/log"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewCore(ctx context.Context, config *C.Config, logger *log.Logger) *Core {
	return &Core{
		config: *config,
		logger: logger,
		ctx:    ctx,
	}
}

func (c *Core) Run() int {
	if c.ctx == nil {
		var cancel context.CancelFunc
		c.ctx, cancel = context.WithCancel(context.Background())
		defer cancel()
	}
	if c.logger == nil {
		c.logger = log.NewLogger(os.Stdout, nil)
	}
	c.logger.Info("Core", "core running...")
	defer c.logger.Info("Core", "core stopped")

	var err error

	// database
	c.logger.Info("Core", "open database...")
	c.logger.Debug("Core", fmt.Sprintf("database driver: %s, url: %s", c.config.DataBase.DriverName, c.config.DataBase.Url))
	c.db = database.NewDatabase(c.ctx)
	err = c.db.Open(c.config.DataBase.DriverName, c.config.DataBase.Url)
	if err != nil {
		c.logger.Fatal("Core", fmt.Sprintf("database open failed: %s", err))
		return 1
	}
	c.db.Logger = c.logger.NewGormLogger()
	defer c.db.Close()
	c.logger.Info("Core", "open database success")
	if c.config.DataBase.Init {
		c.logger.Info("Core", "init database...")
		err = c.db.Init()
		if err != nil {
			c.logger.Fatal("Core", fmt.Sprintf("database init failed: %s", err))
			return 1
		}
		c.logger.Info("Core", "init database success")
	}

	// cache
	c.logger.Info("Core", "open cache...")
	if c.config.Redis.Address == "" {
		c.logger.Info("Core", "use cachemap")
		c.cache = cachemap.New(c.ctx)
	} else {
		c.logger.Info("Core", "use redis")
		c.logger.Debug("Core", fmt.Sprintf("redis address: %s, port: %d, username: %s, password: %s", c.config.Redis.Address, c.config.Redis.Port, c.config.Redis.Username, c.config.Redis.Password))
		c.cache = redis.New(c.ctx, &redis.Options{
			Addr:     net.JoinHostPort(c.config.Redis.Address, strconv.Itoa(int(c.config.Redis.Port))),
			Username: c.config.Redis.Username,
			Password: c.config.Redis.Password,
		})
	}
	defer c.cache.Close()
	c.logger.Info("Core", "open cache success")

	// enable recovery
	defer func() {
		err := recover()
		if err != nil {
			c.logger.Error("Core", fmt.Sprintf("panic: %s", err))
		}
	}()

	// server
	gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.New()
	ginEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: c.logger.NewGinLogger(),
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s: %d %s %s %s %s %s", "Gin Request", params.StatusCode, params.Method, params.Path, params.ClientIP, params.Request.UserAgent(), params.ErrorMessage)
		},
	}))

	// handler
	runningCtx := &handler.RunnginCtx{
		Ctx:    c.ctx,
		Logger: c.logger,
		DB:     c.db,
		Cache:  c.cache,
	}
	c.logger.Info("Core", "init handler...")
	err = runningCtx.Init(ginEngine)
	if err != nil {
		c.logger.Fatal("Core", fmt.Sprintf("handler init failed: %s", err))
		return 1
	}
	defer runningCtx.Close()
	c.logger.Info("Core", "init handler success")

	// run http server
	c.logger.Info("Core", fmt.Sprintf("run http server, listen addr: %s", c.config.ListenAddress))
	server := &http.Server{}
	server.Addr = c.config.ListenAddress
	server.Handler = ginEngine.Handler()
	go func() {
		<-c.ctx.Done()
		_ = server.Shutdown(c.ctx)
	}()
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		c.logger.Fatal("Core", fmt.Sprintf("failed to run http server: %s", err))
		return 1
	} else {
		c.logger.Info("Core", "http server stopped")
		return 0
	}
}
