package service

import (
	"library/common"
	"library/model"
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
		return false
	}
	return true
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
