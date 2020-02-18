package dao

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func (d *UserDao)InsetSessionIdInRedis(sessionId,username string)(err error){
	conn := d.RedisPool.Get()
	if _,err = redis.String(conn.Do("SETEX",sessionId,3600,username));err!=nil{
		return
	}
	return
}
func (d *UserDao)CheckSessionIdInRedis(sessionId string)(username string,err error)  {
	conn := d.RedisPool.Get()
	log.Println(sessionId)
	username, err = redis.String(conn.Do("GET", sessionId))
	if err!=nil{
		log.Println(err)
	}
	return
}