package user

import (
	"net/http"
	"strings"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	var userSignInDto UserSignInDto

	if err := ctx.BindJSON(&userSignInDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	var db = *database.DB
	var user User

	err := db.Where("username = ?", &userSignInDto.UsernameOrEmail).
		Or("email = ?", &userSignInDto.UsernameOrEmail).
		First(&user).
		Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Username or email not registered yet!",
		})

		return
	}

	decryptedPassword := helper.ComparePassword(
		[]byte(user.Password),
		[]byte(userSignInDto.Password),
	)

	if !decryptedPassword {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Wrong password!",
		})

		return
	}

	token := helper.GenerateToken(helper.JwtPayload{
		Email:    user.Email,
		Username: user.Username,
		RoleId:   user.RoleId,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successfully!",
		"data":    token,
	})
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
		})

		return
	}

	if len(users) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found!",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get all user successfully!",
		"data":    users,
	})
}

func DecodeUser(ctx *gin.Context) {
	getHeader := ctx.GetHeader("Authorization")
	authHeader := strings.Split(string(getHeader), " ")
	tokenString := authHeader[1]

	err := helper.VerifyToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized user!",
		})

		return
	}

	user, err := helper.DecodeToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized user!",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Decode user successfully!",
		"data":    user,
	})
}

func GetUserById(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}
