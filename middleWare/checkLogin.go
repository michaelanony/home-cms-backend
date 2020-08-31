package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"home-cms/dao"
)

func CheckLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("auth") !=""{
			fmt.Println(c.Query("auth"))
			c.Next()
		}
		if url := c.Request.URL.String(); url != "/api/user/login" {
			SESSIONID, err := c.Cookie("SESSIONID")
			userNameCookie, err := c.Cookie("username")
			userNameRedis, err := dao.GinDao.CheckSessionIdInRedis(SESSIONID)
			if err != nil || userNameCookie != userNameRedis {
				c.JSON(403, gin.H{
					"code": 403,
					"msg":  "请先登入",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
