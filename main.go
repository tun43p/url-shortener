package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tun43p/api/routes"
)

func main() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&routes.URLResponse{}); err != nil {
		panic("Failed to migrate database")
	}

	e := gin.Default()
	v1 := e.Group("/api/v1")

	v1.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	v1.GET("/urls", func(ctx *gin.Context) {
		routes.GetURLs(ctx, db)
	})

	v1.POST("/urls", func(ctx *gin.Context) {
		routes.ShrinkUrl(ctx, db)
	})

	e.Group("/s").GET("/:s", func(ctx *gin.Context) {
		routes.Redirect(ctx, db)
	})

	e.Run("localhost:8080")
}
