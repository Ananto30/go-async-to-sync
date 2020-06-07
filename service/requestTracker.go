package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

var ResponseMap = make(map[string]gin.H)

var Broadcast = make(chan []byte)
var Response = make(chan io.Reader)

func HandleResponse() {
	b := make([]byte, 1)
	for {
		select {
		case response := <-Response:
			var result map[string]interface{}
			json.NewDecoder(response).Decode(&result)

			fmt.Printf("%+v\n", result)
			ResponseMap[result["trackId"].(string)] = result

			Broadcast <- b
		}
	}

}
