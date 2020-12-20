package service

import (
	"library/common"
	"library/model"
	"library/util"
	"log"
)

func CreateLoan(titleId int, cardNum string) *util.Err {
	borrowerId, err := GetBorrowerIdByCardNum(cardNum)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}

	// 验证用户身份
	if !validateCard(cardNum) {
		log.Println("身份无效！cardNum = ", cardNum)
		return util.Fail("身份无效！")
	}

	// 检查是否有未归还记录
	if HasOverDue(cardNum) {
		log.Println("存在超期记录，借书无效")
		return util.Fail("存在超期记录，借书无效")
	}

	// 检查是否达到借阅最大数量
	if BorrowQuantityIsExceeded(cardNum) {
		log.Println("借书失败！该用户达到最大借书数量！")
		return util.Fail("借书失败！该用户达到最大借书数量！")
	}

	// 检查图书数量
	if GetTotal(titleId) == 0 {
		log.Println("书本数量不足，借书失败！titleId = ", titleId)
		return util.Fail("书本数量不足，借书失败！")
	}

	db := common.GetDB()
	var loan model.Loan
	loan.BorrowerId = borrowerId
	loan.TitleId = titleId
	if err := db.Create(&loan).Error; err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}

	// 图书数量减一
	TotalMinusOne(titleId)

	return util.Success()
}

func GetLoans(cardNum string) (*[]model.Loan, *util.Err) {
	// 验证借阅证是否有效
	if !validateCard(cardNum) {
		log.Println("身份无效！cardNum = ", cardNum)
		return nil, util.Fail("身份无效！")
	}

	// 检查是否有超期未归还图书
	if HasOverDue(cardNum) {
		log.Println("存在超期记录!")
	}

	// 查询借书数量是否达到上限
	if BorrowQuantityIsExceeded(cardNum) {
		log.Println("该用户达到最大借书数量！")
	}

	db := common.GetDB()
	borrowerId, err := GetBorrowerIdByCardNum(cardNum)
	if util.IsFailed(err) {
		log.Println(err.Message)
		return nil, err
	}
	var loans []model.Loan
	if err := db.Where("borrower_id = ?", borrowerId).Find(&loans).Error; err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &loans, util.Success()
}

// 删除一条借阅记录
func DeleteLoan(cardNum string, titleId int) *util.Err {
	borrowerId, err := GetBorrowerIdByCardNum(cardNum)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}

	db := common.GetDB()
	err1 := db.Where("borrower_id = ? AND title_id = ?", borrowerId, titleId).Delete(&model.Loan{}).Error
	if err1 != nil {
		log.Println(err1)
		return util.Fail(err1.Error())
	}
	return util.Success()
}

func GetLoansByTitleId(titleId int) (*[]model.Loan, *util.Err) {
	db := common.GetDB()
	var loans []model.Loan
	err := db.Where("title_id = ?", titleId).Find(&loans).Error
	if err != nil {
		log.Println(err.Error())
		return nil, util.Fail(err.Error())
	}
	return &loans, util.Success()
}
