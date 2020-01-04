package controller

import (
	"github.com/gin-gonic/gin"

	"go-async-to-sync/service"
)

// CallbackHandler will 
func (ctrl AsyncController) CallbackHandler(c *gin.Context) {

	body := c.Request.Body

	service.Response <- body
}
