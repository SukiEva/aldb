package v1

import (
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/SukiEva/aldb/server/util/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		code = e.CODE.FileReceiveError
	} else {
		var fileUrl, fileName string
		cos := upload.NewCos()
		fileUrl, fileName, err = cos.UploadFile(header)
		if err != nil {
			code = e.CODE.FileUploadError
		} else {
			data["name"] = fileName
			data["url"] = fileUrl
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
