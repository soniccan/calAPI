package handlers

import (
	"calAPI/models"
	"calAPI/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListRecipesHandler(c *gin.Context) {

	var reservations []models.DbReservation

	// DBから全てのTODOを取得
	rows, err := utils.Db.
		Query("SELECT * FROM reservations ")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		reservation := models.DbReservation{}
		err = rows.Scan(&reservation.Id, &reservation.UserId, &reservation.ReservationId, &reservation.PaymentID, &reservation.StartAt, &reservation.EndAt)

		if err != nil {
			log.Print(err)
			return
		}
		reservations = append(reservations, reservation)
	}

	c.JSON(http.StatusOK, reservations)
}
