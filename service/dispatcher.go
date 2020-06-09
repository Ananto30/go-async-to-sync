package service

import (
	"fmt"
	"github.com/Ananto30/go-async-to-sync/pubsub"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// SafeMap is mutex managed for our ResMap that holds all the trackId-response mapping
type SafeMap struct {
	ResMap map[string]gin.H
	Length int
	m      sync.Mutex
}

func (i *SafeMap) Get(key string) (gin.H, bool) {
	i.m.Lock()
	defer i.m.Unlock()
	val, found := i.ResMap[key]
	return val, found
}

func (i *SafeMap) Set(key string, val gin.H) {
	i.m.Lock()
	defer i.m.Unlock()
	i.ResMap[key] = val
	i.Length++
}

func (i *SafeMap) Delete(key string) {
	i.m.Lock()
	defer i.m.Unlock()
	delete(ResponseMap.ResMap, key)
	i.Length--
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{}
	sm.ResMap = make(map[string]gin.H)
	sm.Length = 0
	return sm
}

// ResponseMap stores response against trackId
var ResponseMap = NewSafeMap()

// Broadcast channel is responsible to ping waiting clients for a new entry in ResponseMap
var Broadcast = make(chan bool)

// WebhookResponse comes from the callback handler
var WebhookResponse = make(chan gin.H)

// ResponseMapLengthCache is used
var ResponseMapLengthCache = len(ResponseMap.ResMap)

// WebhookDispatcher dispatch a single byte in the Broadcast channel.
// This Broadcast is received by the requests that are hold and listening to Broadcast channel
func WebhookDispatcher() {

	for {
		select {
		case response := <-WebhookResponse:

			//fmt.Printf("Dispatcher %+v\n", response)
			ResponseMap.Set(response["trackId"].(string), response)
			ResponseMapLengthCache++

			Broadcast <- true
		}
	}

}

/*
We can use either of two methods for ensuring all the webhooks are sent to all of the subscribers -
1. Periodically check the ResponseMap
2. Fan-out Broadcast to all subscribers
*/

// PeriodicalMapChecker checks the ResponseMap periodically to make sure all items in the map gets dispatched.
// Downside is when many requests are queued, it will be the slowest to process
func PeriodicalMapChecker() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if ResponseMap.Length > 0 {
				fmt.Println("dispatching from periodical map checker")
				Broadcast <- true
			}
		}
	}

}

var PubSub = pubsub.NewPubsub()
var PubSubCl = PubSubClient{pubsub: PubSub}

//
func WebhookPublisher() {

	for {
		select {
		case response := <-WebhookResponse:

			PubSub.Publish(response["trackId"].(string), response)

		}
	}

}
