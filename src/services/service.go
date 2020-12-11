package services

import (
	"strings"

	"github.com/egnimos/bookstore_oauth_api/src/domain/access_token"
	"github.com/egnimos/bookstore_oauth_api/src/repository/db"
	"github.com/egnimos/bookstore_oauth_api/src/repository/rest"
	"github.com/egnimos/bookstore_oauth_api/src/utils/rest_errors"
)

type ServiceInterface interface {
	GetByID(string) (*access_token.AccessToken, *rest_errors.RestErr)
	CreateAccessToken(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type service struct {
	userRestRepo rest.UsersRepositoryInterface
	dbRepo       db.DBRepository
}

//Create the NewService of given type ServiceInterface and have the repo of type Repository or DBRepository
func NewService(userRestRepo rest.UsersRepositoryInterface, dbRepo db.DBRepository) ServiceInterface {
	return &service{
		userRestRepo: userRestRepo,
		dbRepo:       dbRepo,
	}
}

//GetByID : get the access token by the ID
func (s *service) GetByID(accessTokenID string) (*access_token.AccessToken, *rest_errors.RestErr) {
	//filter the given id
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token")
	}
	//call the getbyid from the dbRepositories
	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

//CreateAccessToken : create access token of the given struct
func (s *service) CreateAccessToken(request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {
	//validate the given access token struct
	if err := request.ValidationRequest(); err != nil {
		return nil, err
	}

	//authenticate the user against the user api
	user, err := s.userRestRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	//generate the access token for the given user
	accessToken := access_token.GetNewAccessToken(user.Id)
	accessToken.Generate()

	//save the new access token in the cassandra
	if err := s.dbRepo.CreateAccessToken(accessToken); err != nil {
		return nil, err
	}

	//call the function from the DBRepositories and return the final answer
	return &accessToken, nil
}

//UpdateExpirationTime : update expiration time of the given access token
func (s *service) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	//validate the give access token
	if err := at.Validation(); err != nil {
		return err
	}
	//call the function from the DBRepositories and return the value or final answer
	return s.dbRepo.UpdateExpirationTime(at)
}
