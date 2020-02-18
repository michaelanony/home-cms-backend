package dao

import (
	"home-cms/model"
	"log"
)

func (d *UserDao)BuyGoodsInDb(goodsName string) (ret model.HomeGoods,err error){
	sqlStr:="SELECT goods_name,goods_store,goods_pay FROM home_cms.home_goods where goods_name = ?"
	if err = d.MysqlPool.Get(&ret,sqlStr,goodsName);err!=nil{
		log.Println(err)
		return
	}
	log.Println(&ret)
	return
}

func (d *UserDao)CutStore(goodsName string) (ret model.HomeGoods,err error){
	sqlStr:="SELECT goods_store FROM home_cms.home_goods where goods_name = ?"
	if err = d.MysqlPool.Get(&ret,sqlStr,goodsName);err!=nil{
		log.Println(err)
		return
	}

	updateStr:="UPDATE home_cms.home_goods SET goods_store = goods_store- 1 WHERE goods_name = ?"
	if _,err = d.MysqlPool.Exec(updateStr, goodsName);err!=nil{
		log.Println(err)
	}
	return
}
