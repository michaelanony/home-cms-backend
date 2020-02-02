package fileHandle

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup)  {
	rr :=r.Group("/file")
	rr.GET("/test", func(c *gin.Context) {
		c.String(200,"test")
	})

}

