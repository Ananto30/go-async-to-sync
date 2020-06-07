package controller

import (
	"net/http"

	"github.com/Ananto30/go-async-to-sync/dto"
	"github.com/Ananto30/go-async-to-sync/service"

	"github.com/gin-gonic/gin"
)

type AsyncController struct{}

func (ctrl AsyncController) GetAsyncInfo(c *gin.Context) {

	var infoReq dto.DemoInfoReq

	if c.BindJSON(&infoReq) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid body", "body": infoReq})
		c.Abort()
		return
	}

	resp, err := service.GetAsyncInfo()
	if err != nil {
		c.Abort()
	}

	c.JSON(http.StatusOK, resp)

}
