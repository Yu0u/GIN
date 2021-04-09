package middleware

//
//import (
//	"gin"
//	"github.com/gin-contrib/cors"
//
//	"time"
//)
//
//func Cors() gin.HandlerFunc {
//	return cors.New(
//		cors.Config{
//			AllowAllOrigins:  true,
//			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//			AllowHeaders:     []string{"*"},
//			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
//			AllowCredentials: true,
//			MaxAge:           12 * time.Hour,
//		},
//	)
//}
