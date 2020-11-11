package main

import (
	"github.com/gin-gonic/gin"
	"ipadgrpc/src/route"
	"log"
)

func main() {
	log.Println("start......")
	ginEngine := gin.New()
	router := new(gin.RouterGroup)
	route.StartApi(router)
	ginEngine.Run(":9999")
}
