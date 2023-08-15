package models

import (
	"time"
)

type Reservation struct {
	UserId    int       `json:"userId"`
	PaymentId int       `json:"paymentId"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
}
type DbReservation struct {
	Id            int       `json:"id"`
	ReservationId int       `json:"reservationId"`
	UserId        int       `json:"userId"`
	PaymentID     int       `json:"paymentId"`
	StartAt       time.Time `json:"startAt"`
	EndAt         time.Time `json:"endAt"`
}

//Name          string    `json:"name"`
//PhoneNumber   string    `json:"phoneNumber"`
