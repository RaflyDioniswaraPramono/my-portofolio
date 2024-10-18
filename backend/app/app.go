package app

import (
	"net/http"

	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	database.DBInit()

	var db = *database.DB

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "API is running well!",
		})
	})

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "API is running well!",
		})
	})

	apiRouterGroup := router.Group("/api")

	{
		user.UserRoutes(&db, apiRouterGroup)
	}

	router.Run(":3000")
}
