package auth

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/tun43p/tun43p.com/internal/api/failure"
)

func AuthMiddleware(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	if ctx.GetHeader("Authorization") != "Bearer "+os.Getenv("API_KEY") {
		ctx.IndentedJSON(401, &failure.FailureResponse{
			Status:  401,
			Message: "Unauthorized",
			Error:   "Invalid API key",
		})

		ctx.Abort()

		return
	}

	ctx.Next()

}
