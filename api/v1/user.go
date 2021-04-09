package v1

import (
	"gin"
	"gin-my/dto"
	"gin-my/middleware"
	"gin-my/model"
	"gin-my/utils/errmsg"
	"gin-my/utils/validator"
	"net/http"
)

var code int

// 注册
func Register(c *gin.Context) {
	var msg string
	var user model.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")

	msg, code = validator.Validate(user)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg,
		})
		c.Abort()
		return
	}

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&user)
	}
	if code == errmsg.USER_EXIST {
		code = errmsg.USER_EXIST
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 登录
func Login(c *gin.Context) {
	var user model.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	var token string
	var code int

	user, code = model.CheckLogin(user.Username, user.Password)

	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "登陆成功",
		"token":   token,
	})
}

// 获取用户信息

func GetInfo(c *gin.Context) {
	user, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"code":    errmsg.SUCCSE,
		"message": errmsg.GetErrMsg(200),
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}
