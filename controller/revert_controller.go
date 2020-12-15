package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

func RevertBook(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	cardNum := mMap["cardNum"]
	titleId, _ := strconv.Atoi(mMap["titleId"])
	if err := service.RevertBook(cardNum, titleId); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "还书成功！")
}
