package router

import (
	"example/chessbot/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	// random ahh cors stuff
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	// router.GET("/searchMove/, handlers.searchMove)
	// router.POST("/bestMove", handlers.)

	router.POST("/search", handlers.Search)
	router.POST("/search_old", handlers.Search_old)

	router.Run(":8080")

}
