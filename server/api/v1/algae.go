package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetData 获取所有数据
func GetData(c *gin.Context) {
	code := e.CODE.Success
	res := model.GetData()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// AddAlga 添加藻类图像
func AddAlga(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var alga model.Alga
	err := c.ShouldBindJSON(&alga)
	if err != nil {
		code = e.CODE.AlgaBindError
	} else {
		res, err := model.AddAlga(alga)
		id := res.(primitive.ObjectID)
		if err != nil {
			code = e.CODE.DataBaseError
		} else {
			if err := model.BindToRiver(alga.River, id); err != nil {
				code = e.CODE.DataBindError
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
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
	code := e.CODE.Success
	res := model.GetRivers()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}
