package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"home-cms/controller"
	"home-cms/dao"
	"home-cms/middleWare"
	"time"
)


func main() {
	if err:= dao.InitPool("michael:cccbbb@tcp(192.168.11.31:30001)/testDb?parseTime=true","192.168.11.31:30002");err!=nil{
		panic(err)
	}
	router :=gin.Default()
	router.Use(middleWare.EnableCookieSession(),cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}),middleWare.CheckLoginStatus())
	rr := controller.GinRouter(router)
	if err:=rr.Run(":80");err!=nil{
		panic(err)
	}
}
