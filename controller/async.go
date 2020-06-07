package controller

import (
	"net/http"

	"github.com/Ananto30/go-async-to-sync/dto"
	"github.com/Ananto30/go-async-to-sync/service"

	"github.com/gin-gonic/gin"
)

// AsyncController ...
type AsyncController struct{}

// GetBalance controller parse the request and send to service
func (ctrl AsyncController) GetBalance(c *gin.Context) {

	var balReq dto.DemoBalanceReq

	if c.BindJSON(&balReq) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid body", "body": balReq})
		c.Abort()
		return
	}

	resp, err := service.GetBalance(balReq.AccountID)
	if err != nil {
		c.Abort()
	}

	c.JSON(http.StatusOK, resp)

}
