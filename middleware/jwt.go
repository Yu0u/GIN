package middleware

import (
	"gin"
	"gin-my/model"
	"gin-my/utils"
	"gin-my/utils/errmsg"
	"github.com/dgrijalva/jwt-go"

	"net/http"
	"strings"

	"time"
)

var JwtKey = []byte(utils.JwtKey)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func SetToken(user model.User) (string, int) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Yu0u",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.INTERNAL_SERVER_ERROR
	}
	return tokenString, errmsg.SUCCSE
}

func ParseToken(token string) (*Claims, int) {
	var claims Claims
	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errmsg.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*Claims); ok && setToken.Valid {
			return key, errmsg.SUCCSE
		} else {
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, pCode := ParseToken(checkToken[1])
		if pCode != errmsg.SUCCSE {
			code = pCode
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		user := model.GetUser(key.Username)
		c.Set("username", user)
		c.Next()
	}
}
