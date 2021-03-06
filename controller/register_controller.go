package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/response"
	"library/service"
	"library/util"
	"strconv"
)

func Register(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	_type, _ := strconv.Atoi(mMap["type"])
	name := mMap["name"]
	cardNum, err := service.Register(name, _type)
	if util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"cardNum": cardNum,
	}, "用户注册成功！")
}
