package bio

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/middlewares"
	"github.com/gin-gonic/gin"
)

func BioRoutes(rg *gin.RouterGroup) {
	rg.POST("/bio", AddBio)
	rg.GET("/bio", middlewares.Auth, GetBios)
	rg.GET("/bio/:id", GetBioById)
	rg.PUT("/bio/:id", UpdateBio)
	rg.DELETE("/bio/:id", DeleteBio)
}
