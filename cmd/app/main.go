package main

import (
	"github.com/tun43p/url-shortener/internal/database"
	"github.com/tun43p/url-shortener/internal/server"
)

func main() {
	server.StartServer(database.CreateDatabase())
}
