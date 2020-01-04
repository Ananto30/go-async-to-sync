package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"go-async-to-sync/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetAsyncInfo() (gin.H, error) {

	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}
	trackID := hex.EncodeToString(u.Bytes())

	demo := &dto.DemoAsyncReq{
		TrackID: trackID,
	}

	// here we call our async service
	resp := MakeRestRequest("http://localhost:5000/try-async", trackID, demo)
	if resp == nil {
		return nil, errors.New("No response from async server")
	}

	return resp, nil
}
