package client

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gowsdl "github.com/hooklift/gowsdl/soap"

	"async-to-sync/service"
)

func MakeRequest(url, action, conversationID string, req, res interface{}) gin.H {

	client := gowsdl.NewClient(url)
	if err := client.Call(action, req, res); err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}

	for {
		resp := <-service.Broadcast
		id, found := resp["conversationId"]
		if found && id == conversationID {
			// TODO: need to delete from the ResponseMap
			return resp
		}

	}

}
