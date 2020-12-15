package service

import (
	"library/util"
	"log"
)

func RevertBook(cardNum string, titleId int) *util.Err {
	// 检查是否为馆藏书
	if !IsInLibrary(titleId) {
		log.Println("还书失败！该书不是本馆藏书！titleId = ", titleId)
		return util.Fail("还书失败！该书不是本馆藏书！")
	}

	// 验证借阅证
	if !validateCard(cardNum) {
		log.Println("借阅证无效！还书失败！")
		return util.Fail("借阅证无效！还书失败！")
	}

	// 判断是否超期
	if IsOverDue(cardNum, titleId) {
		log.Println("还书失败！该记录已经超期！请先进行罚款缴费！")
		return util.Fail("还书失败！该记录已经超期！请先进行罚款缴费！")
	}

	if err := DeleteLoan(cardNum, titleId); util.IsFailed(err) {
		log.Println(err.Message)
		return err
	}
	return util.Success()
}
