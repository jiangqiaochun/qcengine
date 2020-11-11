package route

import (
	"github.com/gin-gonic/gin"
	"ipadgrpc/src/controller"
)

func StartApi(route *gin.RouterGroup) {
	route.POST("/users", func(context *gin.Context) {
		controller.CreateUser(context)
	})
}