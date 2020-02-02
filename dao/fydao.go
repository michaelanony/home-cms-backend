package dao

import (
	"fmt"
	"home-cms/model"
)

//Mysql
//从数据库获取目前
func (d *UserDao)GetFyStaticLatest () (ret *model.StaticNum,err error) {
	ret = &model.StaticNum{}
	sqlStr:="select * from pneumonia_record.static_num order by confirm_num_all desc limit 1"
	err = d.MysqlPool.Get(ret,sqlStr)
	return
}

//获取当前城市的情况
func (d *UserDao)GetCurrentCityStatus (city,param string) (ret []model.CityStatus,err error) {
	ret = make([]model.CityStatus,0)
	fmt.Printf("city is %s,param is %s",city,param)
	tableName :="pneumonia_record."+city
	if param == "all"{
		sqlStr :="select * from "+tableName
		if city =="all"{
			sqlStr ="select * from pneumonia_record.all_city_new"
		}
		fmt.Println(sqlStr)
		err = d.MysqlPool.Select(&ret,sqlStr)
		return
	}else{
		sqlStr:="select * from "+tableName+" order by confirm_num desc limit 1"
		fmt.Println(sqlStr)
		err = d.MysqlPool.Get(&ret,sqlStr)
		return
	}
}

//获取所有城市列表
func (d *UserDao)GetCityListDb() (ret []model.CitySearchPy,err error) {
	ret = make([]model.CitySearchPy,0)
	sqlStr:="select * from pneumonia_record.city_list"
	err = d.MysqlPool.Select(&ret,sqlStr)
	return
}