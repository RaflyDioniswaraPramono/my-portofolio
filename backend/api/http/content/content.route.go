package content

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/middlewares"
	"github.com/gin-gonic/gin"
)

func ContentRoutes(rg *gin.RouterGroup) {
	rg.GET("/content", middlewares.Auth, GetContents)
}
