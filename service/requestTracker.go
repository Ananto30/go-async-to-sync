package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ResponseMap stores response against trackId
var ResponseMap = make(map[string]gin.H)

// Broadcast channel is responsible to ping waiting clients for a new entry in ResponseMap
var Broadcast = make(chan bool)

// WebhookResponse comes from the callback handler
var WebhookResponse = make(chan gin.H)

// WebhookDispatcher dispatch a single byte in the Broadcast channel.
// This Broadcast is recieved by the requests that are hold and listening to Broadcast channel
func WebhookDispatcher() {
	for {
		select {
		case response := <-WebhookResponse:

			fmt.Printf("%+v\n", response)
			ResponseMap[response["trackId"].(string)] = response

			Broadcast <- true
		}
	}

}
