package dao

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func (d *UserDao) RechargeMoney(user string, money int64) (remain int64, err error) {
	redisConn := d.RedisPool.Get()
	ret,err := redis.Int(redisConn.Do("GET",user))
	if ret !=0{
		return
	}
	_, err = redis.String(redisConn.Do("SETEX",user,100,1))
	conn, err := d.MysqlPool.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.Exec("UPDATE home_cms.home_user SET u_money = u_money + ? WHERE u_name = ?", money, user)
	_, err = conn.Exec("INSERT INTO home_cms.order_info (u_name,recharge_money) values (?,?)", user, money)
	if err != nil {
		log.Println(err)
		conn.Rollback()
		return
	}
	conn.Commit()
	sqlStr := "SELECT u_money FROM home_cms.home_user where u_name = ?"
	if err = d.MysqlPool.Get(&remain, sqlStr, user); err != nil {
		log.Println(err)
		return
	}
	return
}
