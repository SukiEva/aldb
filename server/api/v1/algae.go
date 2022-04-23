package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mgo = model.Mgo{}

func GetData(c *gin.Context) {
	algae, err := mgo.QueryAlgae()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, algae)
}
