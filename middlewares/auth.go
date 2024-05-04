package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tun43p/api/common"
)

func AuthMiddleware(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	// Use Bearer token for authorization
	if ctx.GetHeader("Authorization") != "Bearer "+os.Getenv("API_KEY") {
		ctx.IndentedJSON(401, &common.Error{
			Status:  401,
			Message: "Unauthorized",
			Error:   "Invalid API key",
		})

		ctx.Abort()

		return
	}

	ctx.Next()

}
