package main

import (
	"gin/gin"
	"net/http"
)

func main(){
	r := gin.New()
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"tests",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK,"<h1>Hello Gin</h1>")
	})
	r.Run(":9999")
}
