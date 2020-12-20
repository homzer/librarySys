package model

import "time"

type Reserve struct {
	Id          int       `json:"id"`
	BorrowerId  int       `json:"borrower_id"`
	TitleId     int       `json:"title_id"`
	ReserveDate time.Time `json:"reserve_date"`
}
