package service

import (
	"library/common"
	"library/model"
	"library/util"
	"log"
)

func CreateReservation(titleId int, cardNum string) *util.Err {
	// 验证借阅证
	if !validateCard(cardNum) {
		log.Println("身份无效！cardNum = ", cardNum)
		return util.Fail("身份无效！")
	}

	borrowerId, err := GetBorrowerIdByCardNum(cardNum)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}

	db := common.GetDB()
	var reserve model.Reserve
	reserve.BorrowerId = borrowerId
	reserve.TitleId = titleId
	if err := db.Create(&reserve).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}

	return util.Success()
}
