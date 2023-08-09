package handlers

import (
	"calAPI/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListRecipesHandler(c *gin.Context) {

	var reservations []models.Reservation

	// DBから全てのTODOを取得
	rows, err := db.
		Query("SELECT * FROM reservations ")
	if err != nil {
		log.Print(err)
		return
	}

	// １行ごとTODOのEntityにマッピングし、返却用のスライスに追加
	for rows.Next() {
		reservation := models.Reservation{}
		err = rows.Scan(&reservation.ID, &reservation.Name, &reservation.PhoneNumber, &reservation.PaymentID, &reservation.StartAt, &reservation.EndAt)
		if err != nil {
			log.Print(err)
			return
		}
		reservations = append(reservations, reservation)
	}

	c.JSON(http.StatusOK, reservations)
}
