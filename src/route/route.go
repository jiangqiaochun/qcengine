package route

import (
	"github.com/gin-gonic/gin"
	"qcengine/src/controller"
)

func StartApi(route *gin.RouterGroup) {
	route.POST("/users", controller.CreateUser)
	route.DELETE("/users/:id", controller.DeleteUser)
	route.PUT("/users/:id", controller.UpdateUser)
	route.GET("/users/:id", controller.FindUser)
	route.GET("/users", controller.FindUser)
}