package middleWare

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func SessionGenerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,err:=c.Cookie("SESSIONID")
		//err不为空时说明cookie不存在
		if err!=nil{
			nano := time.Now().UnixNano()
			rand.Seed(nano)
			rndNum := rand.Int63()
			sessionId := Md5(Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
			fmt.Println(sessionId)
			//设置cookie
			//第三个选项为过期时间，单位为秒,第四个为所在目录，第五个为domain
			//第六个为是否只能通过http是访问，第七个是否允许别人通过js获取自己的cookie
			c.SetCookie("SESSIONID","test",3600,"/","127.0.0.1",false,false)

		}
		c.Next()
	}
}


func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}