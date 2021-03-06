package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"home-cms/dao"
	"home-cms/errno"
	"home-cms/model"
	"home-cms/tool"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func RegistryHandle(c *gin.Context) {
	ip := c.ClientIP()
	if exist, _ := dao.GinDao.CheckIpInRedis(ip); exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "请不要频繁注册",
		})
		return
	}
	fmt.Println(ip)
	user := &model.HomeUser{}
	user.UName = c.PostForm("u_name")
	user.UPassword = c.PostForm("u_password")
	user.URegisterIp = c.ClientIP()
	if user.UPassword == "" || user.UName == "" {
		fmt.Println(errno.ERROR_USER_FORMAT)
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "格式不正确！",
		})
		return
	}
	name := user.UName
	_, err := dao.GinDao.GetUser(name, "")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "用户已经存在",
		})
		return
	}
	_, err = dao.GinDao.RegistryUser(user)
	if err != nil {
		fmt.Println(err)
	}
	_ = dao.GinDao.InsetIntoRedis(ip)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Register success!",
	})
}

func LoginHandle(c *gin.Context) {
	user := &model.HomeUser{}
	user.UName = c.PostForm("u_name")
	user.UPassword = c.PostForm("u_password")
	if len(c.Request.PostForm) == 1 {
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(data)
		err := json.Unmarshal(data, user)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "账号密码格式不正确！",
			})
			return
		}
	}
	user, err := dao.GinDao.GetUser(user.UName, user.UPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "账号密码错误！",
		})
		return
	}

	if err = tool.SessionGenerator(user.UName, c); err != nil {
		log.Fatal(err)
	}

	//content := map[string]interface{}{"userInfo":user,"userList":user}
	content := map[string]interface{}{"user_name": user.UName}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登入成功！",
		"data": content,
	})
}

func GetAllUser(c *gin.Context) {
	userList, err := dao.GinDao.GetAllUser()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": userList,
	})
}
func CurrentUser(c *gin.Context) {
	sessionId, err := c.Cookie("SESSIONID")
	username, err := dao.GinDao.CheckSessionIdInRedis(sessionId)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登入成功",
		"data": username,
	})
}

func CookieTest(c *gin.Context) {
	//获取客户端是否携带cookie
	cookie, err := c.Cookie("key_cookie")
	//err不为空时说明cookie不存在
	if err != nil {
		fmt.Println(cookie)
		cookie = "NotSet"
		//设置cookie
		//第三个选项为过期时间，单位为秒,第四个为所在目录，第五个为domain
		//第六个为是否只能通过http是访问，第七个是否允许别人通过js获取自己的cookie
		c.SetCookie("key_cookie", "value_cookie", 60, "/", "127.0.0.1", false, false)
	}
	fmt.Println("cookie值为", cookie)
}

//获取用户信息
func UserInfo(c *gin.Context) {
	sessionId, err := c.Cookie("SESSIONID")
	targetUser := c.DefaultQuery("name", "")
	if targetUser == "" {
		c.JSON(201, gin.H{
			"code": 201,
			"msg":  "请输入查询的用户",
		})
		return
	}
	username, err := dao.GinDao.CheckSessionIdInRedis(sessionId)
	ret := &model.HomeUser{}
	if targetUser == username {
		ret, err = dao.GinDao.GetUser(targetUser, "")
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		user, err := dao.GinDao.GetUser(username, "")
		target, err := dao.GinDao.GetUser(targetUser, "")
		if err != nil {
			log.Println(err)
			return
		}
		if user.URole > target.URole {
			ret = target
		}
	}
	if ret.UName != "" {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": ret,
		})
		return
	} else {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "权限不足",
		})
		return
	}
}
func RechargeMoney(c *gin.Context) {
	var username string
	if c.Query("auth") ==""{
		sessionId, _ := c.Cookie("SESSIONID")
		ret, err := dao.GinDao.CheckSessionIdInRedis(sessionId)
		if err != nil {
			return
		}
		username = ret
	}else{
		username = c.Query("auth")
	}
	money, err := strconv.ParseInt(c.PostForm("money"), 10, 64)
	if err != nil || money <= 0 {
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "金额不正确",
		})
		return
	}
	rechargeMoney, err := dao.GinDao.RechargeMoney(username, money)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "金额不正确",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{"remain": rechargeMoney},
	})
}
