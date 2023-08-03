package models

import "time"

type Reservation struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	PaymentID   string    `json:"paymentId"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}
