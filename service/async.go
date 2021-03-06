package service

import (
	"encoding/hex"
	"errors"

	"github.com/Ananto30/go-async-to-sync/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetBalance(accountID string) (gin.H, error) {

	u := uuid.NewV4()
	trackID := hex.EncodeToString(u.Bytes())

	demo := &dto.DemoAsyncReq{
		TrackID: trackID,
		AccountID: accountID,
	}

	// here we call our async service
	//resp := MakeRestRequest("http://localhost:5000/async-balance", trackID, demo)
	resp := PubSubCl.MakeRestRequestSubscriber("http://localhost:5000/async-balance", trackID, demo)
	if resp == nil {
		return nil, errors.New("no response from async server")
	}

	return resp, nil
}
