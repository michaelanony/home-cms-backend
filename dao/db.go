package dao

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)
var (
	GinDao *UserDao
	MysqlPool *sqlx.DB
	RedisPool *redis.Pool
)

type UserDao struct {
	MysqlPool *sqlx.DB
	RedisPool *redis.Pool
}


func NewUserDao(redisPool *redis.Pool,mysqlPool *sqlx.DB) (userDao *UserDao) {
	userDao = &UserDao{
		MysqlPool:mysqlPool,
		RedisPool: redisPool,
	}
	return
}
func InitPool(mysqlDns,RedisDns string) (err error) {
	if err = createMysqlPool(mysqlDns);err!=nil{
		panic(err)
	}
	if err=createRedisPool(RedisDns);err!=nil{
		panic(err)
	}
	GinDao = NewUserDao(RedisPool,MysqlPool)
	return

}
func createMysqlPool(mysqlDns string)(err error) {
	MysqlPool, err = sqlx.Open("mysql", mysqlDns)
	if err!=nil{
		return
	}
	if err=MysqlPool.Ping();err!=nil{
		return
	}
	MysqlPool.SetMaxIdleConns(100)
	MysqlPool.SetMaxOpenConns(16)
	return
}
func createRedisPool(RedisDns string) (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:8,
		MaxActive:0,
		IdleTimeout:100,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp",RedisDns,)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err:=c.Do("PING")
			return err
		},
	}
	return
}