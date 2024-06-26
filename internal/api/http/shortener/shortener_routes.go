package shortener

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShortenerRoutes(e *gin.Engine, rg *gin.RouterGroup, db *gorm.DB) {
	rg.GET("/shortener", func(ctx *gin.Context) {
		GetSingleOrAllShortenedUrls(ctx, db)
	})

	rg.POST("/shortener", func(ctx *gin.Context) {
		ShortenUrl(ctx, db)
	})

	e.Group("/s").GET("/:u", func(ctx *gin.Context) {
		RedirectShortenedUrlToOriginalUrl(ctx, db)
	})
}
