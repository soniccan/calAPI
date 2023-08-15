package handlers

import (
	"calAPI/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListRecipesHandler(c *gin.Context) {

	var reservations []models.DbReservation

	// DBから全てのTODOを取得
	rows, err := db.
		Query("SELECT * FROM reservations ")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		reservation := models.DbReservation{}
		err = rows.Scan(&reservation.Id, &reservation.ReservationId, &reservation.UserId, &reservation.PaymentID, &reservation.StartAt, &reservation.EndAt)

		if err != nil {
			log.Print(err)
			return
		}
		reservations = append(reservations, reservation)
	}

	c.JSON(http.StatusOK, reservations)
}
