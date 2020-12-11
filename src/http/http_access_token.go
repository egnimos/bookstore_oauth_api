package http

import (
	"net/http"

	"github.com/egnimos/bookstore_oauth_api/src/domain/access_token"
	"github.com/egnimos/bookstore_oauth_api/src/services"
	"github.com/egnimos/bookstore_oauth_api/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandlerInterface interface {
	GetByID(*gin.Context)
	CreateAccessToken(*gin.Context)
}

type accessTokenHandler struct {
	service services.ServiceInterface
}

func NewAccessTokenHandler(service services.ServiceInterface) AccessTokenHandlerInterface {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	//get the access token and get the result by passing it to the service function
	accessToken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	//when the access token is been fetched
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) CreateAccessToken(c *gin.Context) {
	//convert it into the struct
	var at access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restError := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	//send the data to the service method create to insert in the database
	accessToken, err := handler.service.CreateAccessToken(at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
