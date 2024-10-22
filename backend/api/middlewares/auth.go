package middlewares

import (
	"net/http"
	"strings"

	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"messgae": "Authorization header required!",
		})

		return
	}

	authHeader := strings.Split(string(header), " ")

	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Authorization header is not valid!",
		})

		return
	}

	tokenString := authHeader[1]

	err := helper.VerifyToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token is not valid!",
		})

		return
	}

	ctx.Next()
}
