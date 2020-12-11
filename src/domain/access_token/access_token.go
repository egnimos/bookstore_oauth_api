package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/egnimos/bookstore_oauth_api/src/utils/crypto_utils"
	"github.com/egnimos/bookstore_oauth_api/src/utils/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

//Request will be in the form of this struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

//ValidationRequest : this method validates the ACCESSTOKENREQUEST struct
func (at *AccessTokenRequest) ValidationRequest() *rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return rest_errors.NewBadRequestError("Invalid grant_type parameter")
	}
	return nil
}

//Response will be in the form of this struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

//Validation : this method validates the ACCESSTOKEN struct
func (at *AccessToken) Validation() *rest_errors.RestErr {
	//validate the access token
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("Invalid access token")
	}
	//validate the userid
	if at.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid UserID")
	}
	// validate the clientid
	if at.ClientID <= 0 {
		return rest_errors.NewBadRequestError("invalid ClientID")
	}
	//validate the expiration time
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

//GetNewAccessToken : this method returns the new ACCESSTOKEN struct after the time expires
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

//check whether the given token is expired or not
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

//Generate : this method generate the new token
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
