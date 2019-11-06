package controller

import (
	"encoding/hex"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"

	"go-async-to-sync/client"
)

type MerchantController struct{}

type MerchantInfoReq struct {
	msisdn string `json:"msisdn"`
}

type CPSreq struct {
	conversationId string `xml:"conversationId"`
}

type CPSack struct {
	ack string `xml:"ack"`
}

type DemoCpsReq struct {
	ConversationId string `json:"conversationId"`
}

func (ctrl MerchantController) GetMerchantInfo(c *gin.Context) {

	var merchantInfo MerchantInfoReq

	if c.BindJSON(&merchantInfo) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid body", "body": merchantInfo})
		c.Abort()
		return
	}

	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		c.Abort()
		return
	}
	cnvID := hex.EncodeToString(u.Bytes())

	// TODO: make exact xml req
	// cpsReq := &CPSreq{
	// 	conversationId: cnvID,
	// }
	// cpsAck := &CPSack{}

	demo := &DemoCpsReq{
		ConversationId: cnvID,
	}

	resp := client.MakeRestRequest("http://localhost:5000/try-async", cnvID, demo)
	if resp == nil {
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)

}
