package main

import (
	"github.com/gin-gonic/gin"

	"async-to-sync/controller"
	"async-to-sync/service"
)

func main() {
	r := gin.Default()

	merchant := new(controller.MerchantController)
	r.POST("/merchantInfo", merchant.GetMerchantInfo)

	go service.HandleResponse()
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}