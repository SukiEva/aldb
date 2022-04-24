package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {
	userEmail := c.Query("email")
	userPwd := c.Query("password")
	code := e.CODE.Success
	data := make(map[string]interface{})
	if model.CheckAuth(userEmail, userPwd) {
		token, err := util.GenToken(userEmail, userPwd)
		if err != nil {
			code = e.CODE.AuthTokenError
		} else {
			data["token"] = token
		}
	} else {
		code = e.CODE.AuthError
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
