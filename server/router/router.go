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
	r.POST("register", v1.AddUser)
	// api
	api := r.Group("api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("data", v1.GetData)
		api.POST("upload", v1.Upload)
	}
	// api/alga
	algae := api.Group("alga")
	{
		algae.POST("add", v1.AddAlga)
		algae.GET("anno", v1.GetAnnotationByAlga)
	}
	// api/anno
	anno := api.Group("anno")
	{
		anno.POST("add", v1.AddAnnotation)
	}
	// api/user
	user := api.Group("user")
	{
		user.GET("info", v1.GetUser)
		user.POST("pwd", v1.ChangePassword)
		user.GET("all", v1.GetUsers)
	}
	// api/river
	river := api.Group("river")
	{
		river.POST("add", v1.AddRiver)
		river.GET("all", v1.GetRivers)
	}
	return r
}
