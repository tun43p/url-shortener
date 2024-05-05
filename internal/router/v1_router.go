package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tun43p/tun43p.com/internal/api/urls"
	"gorm.io/gorm"
)

func V1Router(e *gin.Engine, db *gorm.DB) {
	v1 := e.Group("/api/v1")

	v1.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, "OK")
	})

	v1.GET("/urls", func(ctx *gin.Context) {
		urls.GetURLs(ctx, db)
	})

	v1.POST("/urls", func(ctx *gin.Context) {
		urls.ShrinkUrl(ctx, db)
	})

}
