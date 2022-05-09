package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetAnnotationByAlga(c *gin.Context) {
	code := e.CODE.Success
	alga := c.Query("alga")
	res := model.GetAnnotationByAlga(alga)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

func AddAnnotation(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var anno model.Anno
	err := c.ShouldBindJSON(&anno)
	if err != nil {
		code = e.CODE.AnnoBindError
	} else {
		res, err := model.AddAnnotation(anno)
		id := res.(primitive.ObjectID)
		if err != nil {
			code = e.CODE.DataBaseError
		} else {
			err1 := model.BindToAlga(anno.Alga, id)
			err2 := model.BindToUser(anno.User, id)
			if err1 != nil || err2 != nil {
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
