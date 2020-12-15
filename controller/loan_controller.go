package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/model"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

func LoanBook(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	titleId, _ := strconv.Atoi(mMap["titleId"])
	cardNum := mMap["cardNum"]
	if err := service.CreateLoan(titleId, cardNum); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "借书成功！")
}


func GetLoans(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	cardNum := mMap["cardNum"]
	var loans *[]model.Loan
	var err *util.Err
	if loans, err = service.GetLoans(cardNum); util.IsFailed(err) {
		response.Fail(ctx ,nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"loans": *loans,
	}, "获取借阅信息成功")
}
