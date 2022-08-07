package models

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func InitMysql(datasource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		panic(err)
	}
	if err := engine.Ping(); err != nil {
		log.Println("xorm connect mysql fail")
		return nil
	}
	return engine
}
func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
