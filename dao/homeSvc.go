package dao

import (
	"home-cms/model"
	"log"
)

func (d *UserDao)GetHostsConfig() (ret []model.HostsConfig,err error){
	sqlStr:="SELECT * FROM home_cms.home_hosts_config"
	if err = d.MysqlPool.Select(&ret,sqlStr);err!=nil{
		log.Println(err)
		return
	}
	log.Println(&ret)
	return
}
