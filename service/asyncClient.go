package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	gowsdl "github.com/hooklift/gowsdl/soap"
)

func MakeSoapRequest(url, action, conversationID string, req, res interface{}) gin.H {

	client := gowsdl.NewClient(url)
	if err := client.Call(action, req, res); err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}

	for {
		resp := <-Broadcast
		id, found := resp["conversationId"]
		if found && id == conversationID {
			// TODO: need to delete from the ResponseMap
			return resp
		}

	}

}

func MakeRestRequest(url, conversationID string, body interface{}) gin.H {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("request body:", req.Body)
	bodyR, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bodyR))

	for {
		resp := <-Broadcast
		id, found := resp["conversationId"]
		if found && id == conversationID {
			// TODO: need to delete from the ResponseMap
			return resp
		}

	}

}
