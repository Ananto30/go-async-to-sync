package controller

import (
	"github.com/gin-gonic/gin"

	"go-async-to-sync/service"
)

func (ctrl MerchantController) CallbackHandler(c *gin.Context) {

	// xml body
	body := c.Request.Body
	// json, err := xj.Convert(body)
	// jsonB, err := ioutil.ReadAll(body)

	// if err != nil {
	// 	fmt.Printf("Something went wrong: %s", err)
	// 	c.Abort()
	// 	return
	// }

	service.Response <- body
}
