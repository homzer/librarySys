package service

import (
	"library/model"
	"library/util"
)

// 获取超期记录
func GetOverDues(cardNum string) (*[]model.Loan, *util.Err) {
	return nil, util.Success()
	// 待实现
	// ...
}

// 判读是否有超期记录
func HasOverDue(cardNum string) bool {
	return false
	// 待实现
	// ...
}

// 判断该记录是否超期
func IsOverDue(cardNum string, titleId int) bool {
	return false
	// 待实现
	// ...
}

func GetOverDueFine(cardNum string, titleId int) (float64, *util.Err) {
	// 获取用户类型

	// 获取超期天数

	// 计算结果
	// ...

	return 0, util.Success()
}
