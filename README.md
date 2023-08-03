# calAPI
音楽スタジオを予約するシステムを作りたいと考え、予約をデータベースで管理するAPIを作成している。

DBには、sqlite3を用い、以下の構造体で予約データを保存している。

```golang
type Reservation struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	PaymentID   string    `json:"paymentId"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}
```


# CREATE
`GET /reservations`



# READ
`POST /reserve`


# Todos
UPDATEとDELETEの実装を行う予定である。 

