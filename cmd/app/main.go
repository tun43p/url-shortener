package main

import (
	"github.com/tun43p/tun43p.com/internal/database"
	"github.com/tun43p/tun43p.com/internal/server"
)

func main() {
	server.StartServer(database.CreateDatabase())
}
