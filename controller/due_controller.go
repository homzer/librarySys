package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

// 获取超期罚款金额
func GetOverDueFine(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	titleId, _ := strconv.Atoi(mMap["titleId"])
	cardNum := mMap["cardNum"]
	var dueFine float64
	var err *util.Err
	if dueFine, err = service.GetOverDueFine(cardNum, titleId); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"dueFine": dueFine,
	}, "获取罚款金额成功！")
}
