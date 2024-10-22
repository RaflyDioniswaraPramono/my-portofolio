package bio

import (
	"net/http"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/gin-gonic/gin"
)

func AddBio(ctx *gin.Context) {

}

func GetBios(ctx *gin.Context) {
	var db = *database.DB
	var bios []Bio

	if err := db.Preload("users").Find(&bios).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Can't retrieve biography data from database!",
		})

		return
	}

	if len(bios) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Biography not found!",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get all biography successfully!",
		"data":    bios,
	})
}

func GetBioById(ctx *gin.Context) {

}

func UpdateBio(ctx *gin.Context) {

}

func DeleteBio(ctx *gin.Context) {

}
