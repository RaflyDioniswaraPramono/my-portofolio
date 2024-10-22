package user

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/auth/signin", SignIn)
	rg.POST("/auth/signup", SignUp)
	rg.GET("/auth/decode", middlewares.Auth, DecodeUser)
}

func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/user", middlewares.Auth, GetUsers)
	rg.GET("/user/:id", middlewares.Auth, GetUserById)
	rg.PUT("/user/:id", middlewares.Auth, UpdateUser)
	rg.DELETE("/user/:id", middlewares.Auth, DeleteUser)
}
