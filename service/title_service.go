package service

import (
	"library/common"
	"library/model"
	"library/util"
	"log"
)

func GetTotal(titleId int) int {
	db := common.GetDB()
	var title model.Title
	if err := db.Where("id = ?", titleId).First(&title).Error; err != nil {
		log.Println(err)
		return 0
	}
	return title.Total
}

// 判断是否为本馆藏书
func IsInLibrary(titleId int) bool {
	db := common.GetDB()
	var title model.Title
	if err := db.Where("id = ?", titleId).First(&title).Error; err != nil {
		log.Println(err)
		return false
	}
	if title.Id == 0 {
		log.Println("该图书不存在！", titleId)
		return false
	}
	return true
}

// 判断图书是否在借
func IsBorrowed(titleId int) bool {
	loads, err := GetLoansByTitleId(titleId)
	if util.IsFailed(err) {
		log.Println(err)
		return true
	}
	if len(*loads) > 0 {
		return true
	}
	return false
}

// 书籍在馆数量减一
func TotalMinusOne(titleId int) {
	db := common.GetDB()
	var title model.Title
	if err := db.Where("id = ?", titleId).First(&title).Error; err != nil {
		log.Println(err)
		return
	}
	if title.Total <= 0 {
		log.Println("当前图书数量小于等于0：total = ", title.Total)
		return
	}
	title.Total -= 1
	if err := db.Save(&title).Error; err != nil {
		log.Println(err)
		return
	}
}

// 根据titleId获取title
func GetTitle(titleId int) (*model.Title, *util.Err) {
	db := common.GetDB()
	var title model.Title
	err := db.Where("id = ?", titleId).First(&title).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	if titleId == 0 {
		log.Println("图书不在馆")
		return nil, util.Fail("图书不在馆")
	}
	return &title, util.Success()
}

// 根据titleId获取类型
func GetTitleType(titleId int) (int, *util.Err) {
	// 判断是否在馆
	if !IsInLibrary(titleId) {
		log.Println("该图书不在馆")
		return -1, util.Fail("该图书不在馆")
	}

	title, err := GetTitle(titleId)
	if util.IsFailed(err) {
		log.Println(err)
		return -1, util.Fail(err.Message)
	}
	return title.Type, util.Success()
}

func DeleteTitle(titleId int) *util.Err {

	// 检测图书是否在馆
	if !IsInLibrary(titleId) {
		log.Println("删除失败！该图书不在馆！titleId = ", titleId)
		return util.Fail("删除失败！该图书不在馆！")
	}

	// 检测图书是否在借
	if IsBorrowed(titleId) {
		log.Println("删除失败！该图书在借！titleId = ", titleId)
		return util.Fail("删除失败！该图书在借！")
	}

	return util.Success()
}
