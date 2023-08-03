package handlers

import (
	"calAPI/models"
	"database/sql"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {

	err := os.Remove("./reservations.db")
	if err != nil {
		log.Printf("%q: %s\n", err)
	}

	db, err = sql.Open("sqlite3", "reservations.db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table reservations 
	(id integer not null primary key autoincrement , name text,phoneNumber text,paymentId text,
	StartAt datetime,EndAt datetime);
	delete from reservations;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func AddReservationHandler(c *gin.Context) {
	var reservation models.Reservation
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	cmd := "INSERT INTO reservations (name,phoneNumber,paymentId,startAt,endAt) VALUES (?,?,?,?,?)"
	_, err := db.Exec(cmd, reservation.Name, reservation.PhoneNumber, reservation.PaymentID, reservation.StartAt, reservation.EndAt)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "OK")
}
