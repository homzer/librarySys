package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library/model"
	"strconv"
)

func SignUp(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	var borrower model.Borrower
	borrower.Type, _ = strconv.Atoi(mMap["type"])
	borrower.Name = mMap["name"]

}
