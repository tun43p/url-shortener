package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tun43p/tun43p.com/internal/api/urls"
	"gorm.io/gorm"
)

func RedirectionRouter(e *gin.Engine, db *gorm.DB) {
	e.Group("/s").GET("/:s", func(ctx *gin.Context) {
		urls.Redirect(ctx, db)
	})
}
