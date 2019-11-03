package controller

import (
	"encoding/hex"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"

	"async-to-sync/client"
)

type MerchantController struct{}

type MerchantInfoReq struct {
	msisdn string
}

type CPSreq struct {
	conversationId string `xml:"conversationId"`
}

type CPSack struct {
	ack string `xml:"ack"`
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
	cpsReq := &CPSreq{
		conversationId: cnvID,
	}
	cpsAck := &CPSack{}

	resp := client.MakeRequest("https://soap.example.com/call", "merchantInfoAction", cnvID, cpsReq, cpsAck)
	if resp == nil {
		c.Abort()
		return
	}

}
