package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qcengine/src/model"
	service "qcengine/src/service"
)
var userService = service.NewUserService()

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

func DeleteUser(context *gin.Context) {
	userId := context.Param("id")
	user := new(model.User)
	user.Id = userId
	err := userService.DeleteUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, nil)
	return
}

func UpdateUser(context *gin.Context) {
	user := new(model.User)
	userId := context.Param("id")
	user.Id = userId
	userName := context.PostForm("name")
	if userName != "" {
		user.Name = userName
	}
	err := userService.UpdateUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, nil)
	return
}

func FindUser(context *gin.Context) {
	user := new(model.User)
	userId := context.Param("id")
	if userId == "" {
		userId = context.DefaultQuery("id", "")
	}
	if userId != "" {
		// 根据ID查找
		user.Id = userId
	}
	userName := context.DefaultQuery("name", "")
	if userName != "" {
		user.Name = userName
	}
	res, err := userService.FindUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, res)
	return
}