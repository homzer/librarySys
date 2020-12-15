package service

import (
	"library/common"
	"library/model"
	"library/util"
	"log"
)

func CreateBorrower(borrower *model.Borrower) (*model.Borrower, *util.Err) {
	db := common.GetDB()
	if err := db.Create(&borrower).Error; err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	borrower.CardNum = CreateCardNum(borrower.Id)
	err := db.Save(borrower).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}

	return borrower, util.Success()
}

func GetBorrowerIdByCardNum(cardNum string) (int, *util.Err) {
	db := common.GetDB()
	var borrower model.Borrower
	err := db.Where("card_num = ?", cardNum).First(&borrower).Error
	if err != nil {
		log.Println(err)
		return 0, util.Fail(err.Error())
	}
	return borrower.Id, util.Success()
}

// 获取可借阅最大数量
func GetMaxBorrowingQuantity(cardNum string) int {
	db := common.GetDB()
	var borrower model.Borrower
	err := db.Where("card_num = ?", cardNum).First(&borrower).Error
	if err != nil {
		log.Println(err)
		return 0
	}
	switch borrower.Type {
	case model.GRADUATE:
		return model.QUANTITY2
	case model.DOCTOR:
		return model.QUANTITY3
	case model.UNDERGRADUATE:
		return model.QUANTITY1
	default:
		return model.QUANTITY0
	}
}

// 判断是否超过最大借阅数量
func BorrowQuantityIsExceeded(cardNum string) bool {
	var loans *[]model.Loan
	var err *util.Err
	if loans, err = GetLoans(cardNum); util.IsFailed(err) {
		log.Println(err)
		return true
	}
	max := GetMaxBorrowingQuantity(cardNum)
	if len(*loans) >= max {
		return true
	}
	return false
}
