package routes

import (
	"calAPI/handlers"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// レシピ作成
	router.POST("/reservation", handlers.AddReservationHandler)
	router.GET("/reservations", handlers.ListRecipesHandler)
	router.PUT("/reservation/:id", handlers.UpdateReservationHandler)
	router.DELETE("/reservation/:id", handlers.DeleteReservationHandler)
	return router
}
