package main

import (
	"flag"

	"github.com/alfandevt/go-even-api/internal/database"
	"github.com/alfandevt/go-even-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	runMigration := flag.Bool("migrate", false, "Run Database initialization")
	flag.Parse()

	database.InitDB(*runMigration)
	server := gin.Default()
	routes.RegisterServer(server)
	server.Run(":8000")
}
