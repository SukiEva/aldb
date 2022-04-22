package router

import (
	"github.com/SukiEva/aldb/server/api"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

type AlgaeRouter struct{}

type Group struct {
	AlgaeRouter
}

func init() {
	algaeRouter := Router.Group("api")
	{
		algaeRouter.GET("/data", api.GetData)
	}
}
