package pubsub

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Subs struct {
	ch   chan gin.H
	data gin.H
}

type Pubsub struct {
	mu   sync.RWMutex
	subs map[string]Subs
}

func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string]Subs)
	return ps
}

func (ps *Pubsub) Subscribe(trackID string, ch chan gin.H) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	//ch := make(chan string, 1)

	ps.subs[trackID] = Subs{ch: ch}
}

func (ps *Pubsub) Publish(trackID string, data gin.H) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	//for _, ch := range ps.subs {
	//	go func(ch chan bool) {
	//		ch <- msg
	//	}(ch)
	//}
	ch := ps.subs[trackID].ch
	go func(ch chan gin.H) {
		ch <- data
	}(ch)
}

func (ps *Pubsub) Close(trackID string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch, found := ps.subs[trackID]
	if found {
		close(ch.ch)
		delete(ps.subs, trackID)
	}
}
