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
	// api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("data", v1.GetData)
		api.POST("upload")
	}
	// api/algae
	//algae := api.Group("algae")
	//{
	//	algae.POST("add")
	//}
	// api/user
	//user := api.Group("user")
	//{
	//	user.POST("add", v1.AddUser)
	//}
	// api/river
	river := api.Group("river")
	{
		river.POST("add", v1.AddRiver)
		river.GET("all", v1.GetRivers)
	}
	return r
}
