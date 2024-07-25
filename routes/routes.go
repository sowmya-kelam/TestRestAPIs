package routes

import (
	"restapi/middlewares"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterRoutes(server *gin.Engine, database *sql.DB) {

	server.GET("/events", func(context *gin.Context) {
		getEvents(context, database)
	})

	server.GET("/events/:id", func(context *gin.Context) {
		getEventbyID(context, database)
	})

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", func(context *gin.Context) {
		CreateEvent(context, database)
	})

	authenticated.PUT("/events/:id", func(context *gin.Context) {
		UpdateEvent(context, database)
	})

	authenticated.DELETE("/events/:id", func(context *gin.Context) {
		DeleteEvent(context, database)
	})

	server.POST("/signup", func(context *gin.Context) {
		Signup(context, database)
	})

	server.POST("/login", func(context *gin.Context) {
		Login(context, database)
	})

}
