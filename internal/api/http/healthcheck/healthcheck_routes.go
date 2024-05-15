package healthcheck

import "github.com/gin-gonic/gin"

func HealthcheckRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, "OK")
	})
}
