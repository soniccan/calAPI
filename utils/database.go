package utils

import (
	"calAPI/models"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var Db *sql.DB

func init() {
	err := os.Remove("./reservations.db")
	if err != nil {
		log.Printf("%q: %s\n", err)
	}

	Db, err = sql.Open("sqlite3", "reservations.db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table reservations 
	(ID integer not null primary key autoincrement , USER_ID integer,RESERVATION_ID integer,PAYMENT_ID integer,
	START_AT datetime,END_AT datetime);
	delete from reservations;
	`
	_, err = Db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt = `
	create table users
	(USER_ID integer not null ,PASSWORD text,NAME text,PHONE_NUMBER text);
	delete from users;
	INSERT INTO users (USER_ID,PASSWORD,NAME,PHONE_NUMBER) VALUES (?,?,?,?)
	`

	_, err = Db.Exec(sqlStmt, 22222, "sdsa", "Baba", "0921421421")
	if err != nil {
		log.Fatal(err)
	}
}
func GetAllReservations() ([]models.DbReservation, error) {
	var reservations []models.DbReservation

	// DBから全てのTODOを取得
	rows, err := Db.
		Query("SELECT * FROM reservations ")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		reservation := models.DbReservation{}
		err = rows.Scan(&reservation.Id, &reservation.UserId, &reservation.ReservationId, &reservation.PaymentID, &reservation.StartAt, &reservation.EndAt)

		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}
