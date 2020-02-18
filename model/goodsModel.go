package model

type HomeGoods struct{
	Id	int64	`json:"id" db:"id"`
	GoodsName	string	`json:"goods_name" db:"goods_name"`
	GoodsStore	int64	`json:"goods_store" db:"goods_store"`
	GoodsPay	int64	`json:"goods_pay" db:"goods_pay"`

}
