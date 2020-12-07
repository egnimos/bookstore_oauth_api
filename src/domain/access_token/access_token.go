package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string
	UserID      int64
	ClientID    int64
	Expires     int64
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
