package handlers

import (
	"calAPI/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var db *sql.DB

const ID_RANGE = 10000000

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
	(ID integer not null primary key autoincrement , USER_ID text,RESERVATION_ID integer,PAYMENT_ID integer,
	START_AT datetime,END_AT datetime);
	delete from reservations;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt = `
	create table users
	(USER_ID integer not null ,PASSWORD text,NAME text,PHONE_NUMBER text);
	delete from users;
	INSERT INTO users (USER_ID,PASSWORD,NAME,PHONE_NUMBER) VALUES (?,?,?,?)
	`

	_, err = db.Exec(sqlStmt, 22222, "sdsa", "Baba", "0921421421")
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
	reservationId := rand.Intn(ID_RANGE)

	cmd := "INSERT INTO reservations (USER_ID,RESERVATION_ID,PAYMENT_ID,START_AT,END_AT) VALUES (?,?,?,?,?)"
	_, err := db.Exec(cmd, reservation.UserId, reservationId, reservation.PaymentId, reservation.StartAt, reservation.EndAt)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": reservationId,
	})
	return
}
