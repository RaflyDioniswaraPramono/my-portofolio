package user

import (
	"net/http"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	var userSignInDto UserSignInDto

	if err := ctx.ShouldBindJSON(&userSignInDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
}

func SignUp(ctx *gin.Context) {

}

func GetUsers(ctx *gin.Context) {
	var users []User

	var db = *database.DB

	if err := db.Preload("Role").Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Can't retrieve user data from database!",
			"data":    users,
		})

		return
	}

	if len(users) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found!",
			"data":    users,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get all user successfully!",
		"data":    users,
	})
}

func GetUserById(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}
