package main

import (
	"github.com/gin-gonic/gin"

	"go-async-to-sync/controller"
	"go-async-to-sync/service"
)

func main() {
	r := gin.Default()

	merchant := new(controller.MerchantController)

	r.POST("/merchantInfo", merchant.GetMerchantInfo)
	r.POST("/result", merchant.CallbackHandler)

	go service.HandleResponse()
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}