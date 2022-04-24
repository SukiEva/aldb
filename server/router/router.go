package router

import (
	"github.com/SukiEva/aldb/server/api/v1"
	"github.com/SukiEva/aldb/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.ZapLogger(), middleware.ZapRecovery(true))
	// Auth
	r.GET("/auth", v1.GetAuth)
	// Api Group
	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/data", v1.GetData)
	}
	return r
}
