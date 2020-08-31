package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"home-cms/controller"
	"home-cms/dao"
	"home-cms/middleWare"
	"home-cms/utils"
	"log"
	"time"
)


func main() {
	config, err := utils.InitEnv()
	if err !=nil{
		log.Fatal(err)
	}
	if err:= dao.InitPool(config.DevMysqlDb,config.DevRedis);err!=nil{
		panic(err)
	}
	router :=gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}),middleWare.CheckLoginStatus())
	rr := controller.GinRouter(router)
	if err:=rr.Run(config.Port);err!=nil{
		panic(err)
	}
}
