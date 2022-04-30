package merchant_delivery

import (
	"github.com/gin-gonic/gin"
)

func (res *merchantDelivery) DailyTransaction(c *gin.Context) {
	userId, _ := c.Get("user_id")

	response := res.usecase.DailyTransaction(userId.(int))

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(404, response)
}