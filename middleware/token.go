package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go-crud/api"
	"go-crud/model"
	"time"
)


func TokenCreate() *jwt.GinJWTMiddleware{
	autoToken:=&jwt.GinJWTMiddleware{
		Realm:"test",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour*24,
		MaxRefresh:time.Hour,
		Authenticator:api.AppUserLogin,
		Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code": code,
					"message": message,
				})},
		PayloadFunc:payloadFunc,
	}
	return autoToken
}


func payloadFunc(data interface{})jwt.MapClaims{
	if v,ok:=data.(model.User);ok{
		return jwt.MapClaims{
			"id":v.ID,
			"username":v.UserName,
		}
	}
	return jwt.MapClaims{}
}

