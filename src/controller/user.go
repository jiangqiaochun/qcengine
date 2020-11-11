package controller

import (
	"github.com/gin-gonic/gin"
	"ipadgrpc/src/model"
	"ipadgrpc/src/service"
	"net/http"
)
var userService = servicepackage.NewUserService()

func CreateUser(context *gin.Context) {
	userName := context.PostForm("name")
	user := new(model.User)
	user.Name = userName
	res, err := userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, res)
	return
}