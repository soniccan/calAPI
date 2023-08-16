package handlers

import (
	"calAPI/models"
	"calAPI/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
)

const ID_RANGE = 10000000

func AddReservationHandler(c *gin.Context) {
	var reservation models.Reservation
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	reservationId := rand.Intn(ID_RANGE)

	cmd := "INSERT INTO reservations (USER_ID,RESERVATION_ID,PAYMENT_ID,START_AT,END_AT) VALUES (?,?,?,?,?)"
	_, err := utils.Db.Exec(cmd, reservation.UserId, reservationId, reservation.PaymentId, reservation.StartAt, reservation.EndAt)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": reservationId,
	})
	return
}
