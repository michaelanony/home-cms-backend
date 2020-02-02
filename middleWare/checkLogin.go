package middleWare

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const KEY = "Your secret key"
func EnableCookieSession() gin.HandlerFunc {
	store :=cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID",store)
}
func EnableRedisSession() gin.HandlerFunc {
	store,_:=redis.NewStore(10,"tcp","192.168.11.31:30002","",[]byte(KEY))
	return sessions.Sessions("SESSIONID",store)
}