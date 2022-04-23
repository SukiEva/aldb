package router

import (
	"github.com/SukiEva/aldb/server/api/v1"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

type AlgaeRouter struct{}

type Group struct {
	AlgaeRouter
}

func init() {
	algaeRouter := Router.Group("api/v1")
	{
		algaeRouter.GET("/data", v1.GetData)
	}
}
