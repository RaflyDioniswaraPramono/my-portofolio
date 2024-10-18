package bio

import "github.com/gin-gonic/gin"

func BioRoutes(rg *gin.RouterGroup) {
	rg.POST("/bio", AddBio)
	rg.GET("/bio", GetBios)
	rg.GET("/bio/:id", GetBioById)
	rg.PUT("/bio/:id", UpdateBio)
	rg.DELETE("/bio/:id", DeleteBio)
}
