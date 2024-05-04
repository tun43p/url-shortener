package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tun43p/api/middlewares"
	"github.com/tun43p/api/routes"
)

func main() {
	host := os.Getenv("API_HOST")
	database := os.Getenv("API_DATABASE")

	fmt.Println("Starting API server on", host)

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&routes.URLResponse{}); err != nil {
		panic("Failed to migrate database")
	}

	e := gin.Default()

	v1 := e.Group("/api/v1").Use(middlewares.AuthMiddleware)

	v1.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "OK")
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

	e.Run(host)
}
