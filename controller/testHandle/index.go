package testHandle

import (
	"github.com/gin-gonic/gin"
	"home-cms/service"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/api/test")
	rr.GET("/buy",service.BuyGoods)
}


