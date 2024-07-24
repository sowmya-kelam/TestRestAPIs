package main

import (
	"restapi/db"

	"restapi/routes"

	//"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	defer database.Close()
	server := gin.Default()

	routes.RegisterRoutes(server, database)

	server.Run(":8080")
}
