package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var user model.Operator
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.CODE.UserBindError
	} else if err := model.AddUser(user); err != nil {
		code = e.CODE.UserAlreadyExists
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

func GetUser(c *gin.Context) {
	code := e.CODE.Success
	email := c.Query("email")
	res := model.GetUser(email)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}
