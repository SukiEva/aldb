package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetData 获取所有数据
func GetData(c *gin.Context) {
	res := model.GetData()
	c.JSON(http.StatusOK, res)
}

// AddRiver 添加河流
func AddRiver(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var river model.River
	err := c.ShouldBindJSON(&river)
	if err != nil {
		code = e.CODE.RiverBindError
	} else if err := model.AddRiver(river); err != nil {
		code = e.CODE.RiverAlreadyExists
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetRivers 获取所有河流
func GetRivers(c *gin.Context) {
	res := model.GetRivers()
	c.JSON(http.StatusOK, res)
}
