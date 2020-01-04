package service

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io"
)

var ResponseMap = make(gin.H)

var Broadcast = make(chan gin.H)
var Response = make(chan io.Reader)

func HandleResponse() {

	for {
		select {
		case response := <- Response:
			var result map[string]interface{}
			json.NewDecoder(response).Decode(&result)
			ResponseMap[result["trackId"].(string)] = response

			Broadcast <- result
		}
	}

}