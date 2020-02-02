package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"home-cms/dao"
	"home-cms/model"
)

func GetStaticNumSvc(c *gin.Context){
	sn := &model.StaticNum{}
	sn, err := dao.GinDao.GetFyStaticLatest()
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":sn,
	})
}
func GetCurrentCityStatus(c *gin.Context)  {
	param :=c.DefaultQuery("param","all")
	city :=c.DefaultQuery("city","all")
	cs, err := dao.GinDao.GetCurrentCityStatus(city, param)
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"data":cs,
	})
}

func GetCityList(c *gin.Context)  {
	ret, err := dao.GinDao.GetCityListDb()
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"data":ret,
	})
}