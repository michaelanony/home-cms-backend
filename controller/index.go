package controller

import (
	"github.com/gin-gonic/gin"
	"home-cms/controller/fileHandle"
	"home-cms/controller/fyHandle"
	"home-cms/controller/homeSvcHandle"
	"home-cms/controller/testHandle"
	"home-cms/controller/userHandle"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	rr :=r.Group("/")
	homeSvcHandle.Routers(rr)
	fyHandle.Routers(rr)
	userHandle.Routers(rr)
	fileHandle.Routers(rr)
	testHandle.Routers(rr)
	return r
}
