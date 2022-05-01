package router

import (
	"github.com/SukiEva/aldb/server/api/v1"
	"github.com/SukiEva/aldb/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.ZapLogger(), middleware.ZapRecovery(true))
	r.Use(middleware.Cors())
	// Auth
	r.GET("captcha", v1.GetCaptcha)
	r.POST("auth", v1.GetAuth)
	// api
	api := r.Group("api")
	//api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("data", v1.GetData)
	}
	// api/user
	user := api.Group("user")
	{
		user.POST("add", v1.AddUser)
	}
	return r
}
