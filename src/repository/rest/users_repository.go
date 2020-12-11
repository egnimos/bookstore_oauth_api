package rest

import (
	"encoding/json"
	"time"

	"github.com/egnimos/bookstore_oauth_api/src/domain/users"
	"github.com/egnimos/bookstore_oauth_api/src/utils/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRepositoryInterface interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct{}

func NewRestUserRepository() UsersRepositoryInterface {
	return &usersRepository{}
}

func (ur *usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	//get the response
	response := usersRestClient.Post("users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewBadRequestError("invalid restclient response when trying to login the user")
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface while login the user")
		}
		return nil, apiErr
	}

	//if there is no error then validate the user
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users while login")
	}
	return &user, nil
}
