package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/auth/signin", SignIn)
	rg.POST("/auth/signup", SignUp)
}

func UserRoutes(db *gorm.DB, rg *gin.RouterGroup) {
	rg.GET("/user", GetUsers)
	rg.GET("/user/:id", GetUserById)
	rg.PUT("/user/:id", UpdateUser)
	rg.DELETE("/user/:id", DeleteUser)
}
