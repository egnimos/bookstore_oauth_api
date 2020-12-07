package access_token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//TestExpirationTime
func TestExpirationTime(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be of 24 hours")
}

//TestGetNewAccessToken
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	//New Access Token should not be expires...
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	//access toke should be empty
	assert.Empty(t, at.AccessToken, "new access token should not have been defined yet")
	//userid should be empty
	assert.EqualValues(t, 0, at.UserID, "new access token should not have an userid associated with it")
}

//TestAccessTokenIsExpires
func TestAccessTokenIsExpires(t *testing.T) {
	at := AccessToken{}
	fmt.Println(at.Expires)
	fmt.Println(at.IsExpired())
	assert.True(t, at.IsExpired(), "brand new access token should not be expired")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	fmt.Println(at.Expires)
	fmt.Println(at.IsExpired())
	assert.False(t, at.IsExpired(), "access token should be expires 3 hours from now")
	// if at.IsExpired() {
	// 	t.Error("access token should be expires 3 hours from now")
	// }
}
