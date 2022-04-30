package user_delivery

import (
	"fmt"
	"majoo/helpers"
	"majoo/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (res *userDelivery) Login(c *gin.Context) {
	request := dto.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.SendResponse(http.StatusBadRequest, "Bad Request", err, nil)
		fmt.Printf("%+v", errorRes)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.Login(request)

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(404, response)
}

func (res *userDelivery) Register(c *gin.Context) {
	request := dto.RegisterRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.SendResponse(http.StatusBadRequest, "Bad Request", err, nil)
		fmt.Printf("%+v", errorRes)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.Register(request)

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(404, response)
}