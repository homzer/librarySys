package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

func ReserveBook(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	titleId, _ := strconv.Atoi(mMap["titleId"])
	cardNum := mMap["cardNum"]
	if err := service.CreateReservation(titleId, cardNum); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}

	response.Success(ctx, nil, "结束成功")
}
