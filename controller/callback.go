package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"async-to-sync/service"

	xj "github.com/basgys/goxml2json"
)

func (ctrl MerchantController) callbackController(c *gin.Context) {

	body := c.Request.Body

	json, err := xj.Convert(body)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		c.Abort()
		return
	}

	service.Response <- json
}
