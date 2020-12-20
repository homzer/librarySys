package model

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type Magazine struct {
	Id        int       `json:"id"`
	Printer   string    `json:"printer"`
	PrintDate date.Date `json:"print_date"`
}
