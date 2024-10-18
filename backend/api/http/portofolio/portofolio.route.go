package portofolio

import "github.com/gin-gonic/gin"

func PortofolioRoutes(rg *gin.RouterGroup) {
	rg.POST("/portofolio", AddPortofolio)
	rg.GET("/portofolio", GetPortofolios)
	rg.GET("/portofolio/:id", GetPortofolioById)
	rg.PUT("/portofolio/:id", UpdatePortofolio)
	rg.DELETE("/portofolio/:id", DeletePortofolio)
}
