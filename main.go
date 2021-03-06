package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Ananto30/go-async-to-sync/controller"
	"github.com/Ananto30/go-async-to-sync/service"
)

func main() {
	r := gin.Default()

	asyncCtrl := new(controller.AsyncController)

	r.POST("/async", asyncCtrl.GetBalance)
	r.POST("/result", asyncCtrl.CallbackHandler)

	// Start the dispatcher for webhook callbacks
	// go service.WebhookDispatcher()
	// go service.PeriodicalMapChecker()
	go service.WebhookPublisher()

	r.Run(":8005") // listen and serve on 0.0.0.0:8005 (for windows "localhost:8005")
}
