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
	r.POST("upload", v1.Upload)
	// api
	api := r.Group("api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("data", v1.GetData)
	}
	// api/alga
	algae := api.Group("alga")
	{
		algae.GET("anno", v1.GetAnnotationByAlga)
		algae.POST("add", v1.AddAlga)
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
		user.GET("anno", v1.GetAnnotationByUser)
		user.GET("all", v1.GetUsers)
		user.GET("delete", v1.DeleteUser)
		user.POST("pwd", v1.ChangePassword)
		user.POST("update", v1.UpdateUser)
	}
	// api/river
	river := api.Group("river")
	{
		river.GET("all", v1.GetRivers)
		river.POST("add", v1.AddRiver)
	}
	return r
}
