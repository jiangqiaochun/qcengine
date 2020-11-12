package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"qcengine/src/route"
)

func main() {
	log.Println("start......")
	ginEngine := gin.New()
	v1Group := ginEngine.Group("v1")
	route.StartApi(v1Group)
	ginEngine.Run(":9999")
}
