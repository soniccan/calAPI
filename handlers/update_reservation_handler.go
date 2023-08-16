package handlers

import (
	"calAPI/models"
	"calAPI/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UpdateReservationHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var reservation models.Reservation

	if err := c.ShouldBindJSON(&reservation); err != nil {
		// エラーの場合は、400エラーを返す
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	reservations, err := utils.GetAllReservations()

	if err != nil {
		log.Print(err)
		return
	}
	targetReservationIndex := -1
	for i := 0; i < len(reservations); i++ {
		if reservations[i].ReservationId == id {
			targetReservationIndex = reservations[i].Id

		}
	}

	// もし更新するレシピがない場合
	if targetReservationIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Reservation not found"})
		return
	}

	// 更新する
	sqlStmt := `UPDATE reservations  SET USER_ID=?, PAYMENT_ID=?, START_AT=? , END_AT=? WHERE ID=?`

	_, err = utils.Db.Exec(sqlStmt, reservation.UserId, reservation.PaymentId, reservation.StartAt, reservation.EndAt, targetReservationIndex)

	if err != nil {
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, reservation)

}
