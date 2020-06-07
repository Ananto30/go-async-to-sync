package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Ananto30/go-async-to-sync/service"
)

// CallbackHandler handles the webhook callback from the async server
func (ctrl AsyncController) CallbackHandler(c *gin.Context) {

	body := c.Request.Body

	var jsnBdy map[string]interface{}
	err := json.NewDecoder(body).Decode(&jsnBdy)
	if err != nil {
		fmt.Println("Callback conversion failed", err)
	}

	service.WebhookResponse <- jsnBdy
}
