package service

import (
	"library/common"
	"library/model"
	"log"
)

func CreateCardNum(borrowerId int) string {
	return string(borrowerId + 12345)
}

func validateCard(cardNum string) bool {
	db := common.GetDB()
	var card model.Card
	if err := db.Where("card_num = ?", cardNum).First(&card).Error; err != nil {
		log.Println(err)
		return false
	}
	if len(card.CardNum) == 0 {
		return false
	}
	return true
}
