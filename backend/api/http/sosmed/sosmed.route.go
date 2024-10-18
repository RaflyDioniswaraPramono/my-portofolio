package sosmed

import "github.com/gin-gonic/gin"

func SosmedRoutes(rg *gin.RouterGroup) {
	rg.POST("/sosmed", AddSosmed)
	rg.GET("/sosmed", GetSosmeds)
	rg.GET("/sosmed/:id", GetSosmedById)
	rg.PUT("/sosmed/:id", UpdateSosmed)
	rg.DELETE("/sosmed/:id", DeleteSosmed)
}
