package service

import (
	"github.com/gin-gonic/gin"
	"home-cms/dao"
	"log"
)

func BuyGoods(c *gin.Context)  {
	goodsName := c.DefaultQuery("name","jack")
	goods,err:=dao.GinDao.BuyGoodsInDb(goodsName)
	if err!=nil{
		log.Print(err)
		return
	}
	if goods.GoodsStore<=0{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"failed!",
		})
	}
	store, err := dao.GinDao.CutStore(goodsName)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"data":store,
	})
}