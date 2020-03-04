package homeSvcHandle
import (
	"github.com/gin-gonic/gin"
	"home-cms/service"
)

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/home")
	rr.GET("/hostsconfig",service.GetHomeHosts)
	rr.GET("/twrpcheck",)

}

