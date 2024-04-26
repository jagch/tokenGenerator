package delivery

import (
	"jagch/tokenGenerator/application"
	"jagch/tokenGenerator/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	tokenService application.TokenService
}

func NewTokennController(tokenService application.TokenService) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

func (ctrl *TokenController) GenerateTokens(c *gin.Context) {
	quantityString := c.Param("quantity")
	whitelabelName := c.Param("whitelabelName")

	quantity, err := strconv.ParseUint(quantityString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  -1,
			Message: "Param quantity is incorrect",
			Data:    nil,
		})
		return
	}

	tokens, err := ctrl.tokenService.GenerateTokens(whitelabelName, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  -1,
			Message: "Error generating the tokens",
			Data:    nil,
		})

		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  0,
		Message: "Tokens successfully created",
		Data:    tokens,
	})
}

func (ctrl *TokenController) CheckToken(c *gin.Context) {
	token := c.Param("token")
	whitelabelName := c.Param("whitelabelName")

	exists := ctrl.tokenService.CheckToken(token, whitelabelName)
	if exists {
		c.JSON(http.StatusOK, model.Response{
			Status:  0,
			Message: "The token exists",
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  -1,
			Message: "The token doesn't exists",
			Data:    nil,
		})
	}
}
