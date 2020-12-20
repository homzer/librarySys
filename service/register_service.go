package service

import (
	"library/common"
	"library/model"
	"library/util"
	"log"
)

func Register(name string, _type int) (string, *util.Err) {
	// 创建用户
	var borrower *model.Borrower
	var err *util.Err
	if borrower, err = CreateBorrower(name, _type); util.IsFailed(err) {
		log.Println(err)
		return "", err
	}

	// 创建卡号
	borrower.CardNum = CreateCardNum(borrower.Id)
	db := common.GetDB()
	err1 := db.Save(borrower).Error
	if err1 != nil {
		log.Println(err1)
		return "", util.Fail(err1.Error())
	}

	// 返回卡号
	return borrower.CardNum, util.Success()
}
