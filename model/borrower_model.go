package model

type Borrower struct {
	Id      	int    `json:"id"`
	Name        string `json:"name"`
	CardNum 	string `json:"card_num"`
	Type   	 	int `json:"type"`
}

const (
	GRADUATE  = 0
	UNDERGRADUATE = 1
	DOCTOR = 2
	COLLEGESTUDENT = 3
	TEACHER = 4

	// 可借阅数量
	QUANTITY0 = 3
	QUANTITY1 = 5
	QUANTITY2 = 6
	QUANTITY3 = 10
)
