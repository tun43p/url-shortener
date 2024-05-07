package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tun43p/url-shortener/internal/router"
)

func StartServer(db *gorm.DB) {
	host := os.Getenv("API_HOST")

	e := gin.Default()
	e.SetTrustedProxies(nil)

	router.V1Router(e, db)

	fmt.Println("Starting API server on", host)

	e.Run(host)
}
