package model

type Title struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Type  int     `json:"type"`
	Total int     `json:"total"`
}

const (
	PAPER    = 0
	MAGAZINE = 1
	BOOK     = 2
)
