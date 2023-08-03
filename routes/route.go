package routes

import (
	"calAPI/handlers"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// レシピ作成
	router.POST("/reserve", handlers.AddReservationHandler)
	router.GET("/reservations", handlers.ListRecipesHandler)
	return router
}
