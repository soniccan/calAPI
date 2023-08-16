package handlers

import (
	"calAPI/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteReservationHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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

	sqlStmt := `DELETE FROM reservations WHERE ID=?`

	_, err = utils.Db.Exec(sqlStmt, targetReservationIndex)

	if err != nil {
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
