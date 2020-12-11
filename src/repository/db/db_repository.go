package db

import (
	"errors"
	"fmt"

	"github.com/egnimos/bookstore_oauth_api/src/clients/cassandra"
	"github.com/egnimos/bookstore_oauth_api/src/domain/access_token"
	"github.com/egnimos/bookstore_oauth_api/src/utils/rest_errors"
	"github.com/gocql/gocql"
)

//CASSADRA QUERIES
const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpiresTime = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

//create a new instance of dbRepository
func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *rest_errors.RestErr)
	CreateAccessToken(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {
}

//GetByID : get the entity by the ID
func (db *dbRepository) GetByID(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("there is no access_token with the given ID")
		}
		return nil, rest_errors.NewInternalServerError("database connection is not implemented yet")
	}
	// get access token from the cassandra DB
	return &result, nil
}

//CreateAccessToken : create a accessToken in the database
func (db *dbRepository) CreateAccessToken(at access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintln("error when trying to save the access token in database", err))
	}
	return nil
}

//UpdateExpirationTime : updates the expiration time
func (db *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryUpdateExpiresTime,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintln("error when trying to update the current resource", errors.New("database error")))
	}
	return nil
}
