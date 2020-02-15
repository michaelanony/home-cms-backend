package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"home-cms/dao"
	"home-cms/errno"
	"home-cms/model"
	"io/ioutil"
	"log"
	"net/http"
)

func RegistryHandle(c *gin.Context)  {
	ip := c.ClientIP()
	if exist,_ :=dao.GinDao.CheckIpInRedis(ip);exist{
		c.JSON(http.StatusOK,gin.H{
			"code":300,
			"msg":"请不要频繁注册",
		})
		return
	}
	fmt.Println(ip)
	user :=&model.HomeUser{}
	user.UName=c.PostForm("u_name")
	user.UPassword=c.PostForm("u_password")
	user.URegisterIp=c.ClientIP()
	if user.UPassword==""|| user.UName == "" {
		fmt.Println(errno.ERROR_USER_FORMAT)
		c.JSON(http.StatusOK,gin.H{
			"code":300,
			"msg":"格式不正确！",
		})
		return
	}
	name := user.UName
	_, err := dao.GinDao.GetUser(name,"")
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":300,
			"msg":"用户已经存在",
		})
		return
	}
	_ ,err = dao.GinDao.RegistryUser(user);
	if err!=nil{
		fmt.Println(err)
	}
	_ = dao.GinDao.InsetIntoRedis(ip)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"Register success!",
	})
}

func LoginHandle(c *gin.Context)  {
	user :=&model.HomeUser{}
	user.UName=c.PostForm("u_name")
	user.UPassword=c.PostForm("u_password")
	if len(c.Request.PostForm) == 1{
		data,_:=ioutil.ReadAll(c.Request.Body)
		fmt.Println(data)
		err := json.Unmarshal(data, user)
		if err!=nil{
			fmt.Println(err)
			c.JSON(http.StatusOK,gin.H{
				"code":401,
				"msg":"账号密码格式不正确！",
			})
			return
		}
	}
	user ,err:= dao.GinDao.GetUser(user.UName,user.UPassword)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":401,
			"msg":"账号密码错误！",
		})
		return
	}
	session:=sessions.Default(c)
	session.Set("username",user.UName)
	session.Save()
	content := map[string]interface{}{"userInfo":user,"userList":user}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"登入成功！",
		"data":content,
	})
}

func GetAllUser(c *gin.Context)  {
	userList, err := dao.GinDao.GetAllUser()
	if err!=nil{
		fmt.Println(err)
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":userList,
	})
}
func CurrentUser(c *gin.Context)  {
	log.Println(c.Request.Cookies())
	session := sessions.Default(c)
	username,ok:= session.Get("username").(string)
	if !ok{
		c.JSON(403,gin.H{
			"code":403,
			"msg":"请先登入",
			"data":"",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"登入成功",
		"data":username,
	})
}
