package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Ananto30/go-async-to-sync/controller"
	"github.com/Ananto30/go-async-to-sync/service"
)

func main() {
	r := gin.Default()

	asyncCtrl := new(controller.AsyncController)

	r.POST("/async", asyncCtrl.GetAsyncInfo)
	r.POST("/result", asyncCtrl.CallbackHandler)

	go service.HandleResponse()

	r.Run(":8005") // listen and serve on 0.0.0.0:8005 (for windows "localhost:8005")
}
