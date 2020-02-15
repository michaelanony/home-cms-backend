package userHandle

import (
	"github.com/gin-gonic/gin"
	"home-cms/service"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/api/user")
	rr.POST("/login", service.LoginHandle)
	rr.POST("/registry", service.RegistryHandle)
	rr.GET("/current",service.CurrentUser)
	rr.GET("/alluser",service.GetAllUser)
}


