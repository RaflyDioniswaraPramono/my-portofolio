package content

import (
	"net/http"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/gin-gonic/gin"
)

func GetContents(ctx *gin.Context) {
	var contents []Content

	var db = *database.DB

	if err := db.Find(&contents).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Can't retrieve content data from database!",
			"data":    contents,
		})

		return
	}

	if len(contents) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Content not found!",
			"data":    contents,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"messgae": "Get all content successfully!",
		"data":    contents,
	})
}
