package model

import (
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

type LostTitle struct {
	Id         int       `json:"id"`
	BorrowerId int       `json:"borrower_id"`
	TitleId    int       `json:"title_id"`
	CreatedAt  time.Time `json:"created_at"`
	DueDate    date.Date `json:"due_date"`
}
