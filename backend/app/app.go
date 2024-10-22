package app

import (
	"net/http"
	"time"

	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/content"
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
	"github.com/RaflyDioniswaraPramono/my-portofolio/configs/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	database.DBInit()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your allowed origins here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		user.AuthRoutes(apiRouterGroup)
		user.UserRoutes(apiRouterGroup)
		content.ContentRoutes(apiRouterGroup)
	}

	router.Run(":3000")
}
