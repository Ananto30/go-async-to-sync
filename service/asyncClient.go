package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ananto30/go-async-to-sync/pubsub"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MakeSoapRequest for SOAP request, legacy services
//func MakeSoapRequest(url, action, trackID string, req, res interface{}) gin.H {
//
//	client := gowsdl.NewClient(url)
//	if err := client.Call(action, req, res); err != nil {
//		fmt.Printf("Something went wrong: %s", err)
//		return nil
//	}
//
//	for {
//		<-Broadcast
//		asyncResp, found := ResponseMap[trackID]
//		if found {
//			// TODO: need to delete from the ResponseMap
//			return asyncResp
//		}
//
//	}
//
//}

// MakeRestRequest makes a REST request and wait for a signal from Broadcast
func MakeRestRequest(url, trackID string, body interface{}) gin.H {

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
	bodyR, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bodyR))

	for {
		<-Broadcast
		asyncResp, found := ResponseMap.Get(trackID)

		fmt.Println("paisi mama", asyncResp, found, trackID, ResponseMap)
		if found {
			ResponseMap.Delete(trackID)
			return asyncResp
		}

	}

}

type PubSubClient struct {
	pubsub *pubsub.Pubsub
}

func (ps *PubSubClient) MakeRestRequestSubscriber(url, trackID string, body interface{}) gin.H {

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
	bodyR, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bodyR))

	ch := make(chan gin.H)

	ps.pubsub.Subscribe(trackID, ch)

	data := <-ch

	ps.pubsub.Close(trackID)

	return data

}
