package middleWare

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func CheckLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if url :=c.Request.URL.String();url!="/api/user/login"{
			session := sessions.Default(c)
			_,ok:= session.Get("username").(string)
			if !ok{
				c.JSON(403,gin.H{
					"code":403,
					"msg":"请先登入",
					"data":"",
				})
				c.Abort()
			}
		}
		c.Next()
	}
}