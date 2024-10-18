package experience

import "github.com/gin-gonic/gin"

func ExperienceRoutes(rg *gin.RouterGroup) {
	rg.POST("/experience", AddExperience)
	rg.GET("/experience", GetExperiences)
	rg.GET("/experience/:id", GetExperienceById)
	rg.PUT("/experience/:id", UpdateExperience)
	rg.DELETE("/experience/:id", DeleteExperience)
}
