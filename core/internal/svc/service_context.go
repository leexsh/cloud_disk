package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"leexsh/file_sys/core/internal/config"
	"leexsh/file_sys/core/internal/middleware"
	"leexsh/file_sys/core/models"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config      config.Config
	Engine      *xorm.Engine
	RedisEngine *redis.Client
	Auth        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Engine:      models.InitMysql(c.Mysql.DataSource),
		RedisEngine: models.InitRedis(c.Redis.Addr),
		Auth:        middleware.NewAuthMiddleware().Handle,
	}
}
