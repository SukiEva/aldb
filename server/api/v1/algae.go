package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(c *gin.Context) {
	data := model.GetData()
	c.JSON(http.StatusOK, data)
}
