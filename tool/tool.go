package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"home-cms/dao"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)
var cookieDomain =".cms.home"

func SessionGenerator(username string,c *gin.Context) (err error) {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	sessionId := Md5(Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
	if err =dao.GinDao.InsetSessionIdInRedis(sessionId,username);err!=nil{
		log.Println(err)
		return
	}
	//设置cookie
	//第三个选项为过期时间，单位为秒,第四个为所在目录，第五个为domain
	//第六个为是否只能通过http是访问，第七个是否允许别人通过js获取自己的cookie
	c.SetCookie("SESSIONID",sessionId,3600,"/",cookieDomain,false,false)
	c.SetCookie("username",username,3600,"/",cookieDomain,false,false)
	return
}


func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}