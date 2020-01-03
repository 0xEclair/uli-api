package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	//config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:8080", "http://www.yoshino.studio:3389","http://101.132.34.156:3389","http://yoshino.studio:3389","http://192.168.10.120:3000", "http://192.168.10.120:8080"}


	config.AllowCredentials = true
	return cors.New(config)
}
