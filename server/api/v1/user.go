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

func ChangePassword(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var ch change
	err := c.ShouldBindJSON(&ch)
	if err != nil {
		code = e.CODE.UserBindError
	} else {
		if model.CheckAuth(ch.Email, ch.Password) {
			if err := model.ChangePassword(ch.Email, ch.NewPassword); err != nil {
				code = e.CODE.DataBaseError
			}
		} else {
			code = e.CODE.AuthError
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

type change struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
