package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tun43p/tun43p.com/internal/api/auth"
	"github.com/tun43p/tun43p.com/internal/api/http/healthcheck"
	"github.com/tun43p/tun43p.com/internal/api/http/shortener"
)

func V1Router(e *gin.Engine, db *gorm.DB) {
	v1 := e.Group("/api/v1")
	v1.Use(auth.AuthMiddleware)

	healthcheck.HealthcheckRoutes(v1)
	shortener.ShortenerRoutes(e, v1, db)
}
