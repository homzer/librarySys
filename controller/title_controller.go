package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

func DeleteTitle(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	titleId, _ := strconv.Atoi(mMap["titleId"])
	if err := service.DeleteTitle(titleId); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除标题信息成功！")
}
