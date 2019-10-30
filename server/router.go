package server

import (
	"go-crud/api"
	"go-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	appmidWare:=middleware.TokenCreate()
	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed:=v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)

			authed.POST("videos",api.CreateVideo)
			authed.PUT("video/:id",api.UpdateVideo)
			//	v1.DELETE("video/:id",api.DeleteVideo)
		}
		// app用户登录
		v1.POST("app/user/login",appmidWare.LoginHandler)
		// refresh接口 不会验证token是否完整
		//v1.POST("app/user/refresh",appmidWare.RefreshHandler)
		appAuthed:=v1.Group("/app")
		appAuthed.Use(appmidWare.MiddlewareFunc())
		{
			appAuthed.GET("user/me",api.UserMe)
			appAuthed.GET("user/ping",api.AppPing)
		}



		v1.GET("video/:id",api.ShowVideo)
		v1.GET("videos",api.ListVideo)

		// rank
		v1.GET("rank/daily",api.DailyRank)

		v1.POST("upload/token",api.UploadToken)
	}
	return r
}
