package routers

import (
	"gin"
	v1 "gin-my/api/v1"
	"gin-my/middleware"
	"gin-my/utils"
)

func InitRouter() {

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	//r.Use(middleware.Cors())

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("/user/info",v1.GetInfo)
	}

	r.POST("/register",v1.Register)
	r.POST("/login",v1.Login)

	_ = r.Run(utils.HttpPort)
}
