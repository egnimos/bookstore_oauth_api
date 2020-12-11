package app

import (
	"fmt"

	"github.com/egnimos/bookstore_oauth_api/src/http"
	"github.com/egnimos/bookstore_oauth_api/src/repository/db"
	"github.com/egnimos/bookstore_oauth_api/src/repository/rest"
	"github.com/egnimos/bookstore_oauth_api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(services.NewService(rest.NewRestUserRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.CreateAccessToken)

	fmt.Println("Server is Started....")
	router.Run(":8080")
}
