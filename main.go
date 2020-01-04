package main

import (
	"github.com/gin-gonic/gin"

	"go-async-to-sync/controller"
	"go-async-to-sync/service"
)

func main() {
	r := gin.Default()

	asyncCtrl := new(controller.AsyncController)

	r.POST("/async", asyncCtrl.GetAsyncInfo)
	r.POST("/result", asyncCtrl.CallbackHandler)

	go service.HandleResponse()
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}